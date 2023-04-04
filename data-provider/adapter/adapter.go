package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/bootstrap"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/protoadapter"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type DataProvider struct {
	protoadapter.UnimplementedDataProviderServiceServer
	bootstraps    map[string]*protonode.DPInfo // A list of 3 default bootstrap dp nodes
	knownPeers    map[string]*protonode.DPInfo // A list of dp nodes that are not bootstrap nodes, and only known to bootstrap nodes
	mu            sync.Mutex
	listenAddress string
}

func NewDataProvider() *DataProvider {
	bootstraps := make(map[string]*protonode.DPInfo)
	bootstrapNodes := bootstrap.Nodes()

	for i, dp := range bootstrapNodes {
		bootstraps[dp] = &protonode.DPInfo{
			Name:          fmt.Sprintf("Bootstrap-00%d", i),
			ListenAddress: dp,
			KOrgId:        fmt.Sprintf("bootstrap-00%d", i),
		}
	}

	return &DataProvider{
		knownPeers:    make(map[string]*protonode.DPInfo),
		bootstraps:    bootstraps,
		listenAddress: os.Getenv("HOST_IP"),
	}
}

func NewAdapter() *protoadapter.Adapter {
	return &protoadapter.Adapter{}
}

func ListAdapters() []*protoadapter.Adapter {
	var adapters []*protoadapter.Adapter

	for _, adp := range config.Loaded.Feed.Adapters {
		adapterPath := path.Join(config.Loaded.RootWD, config.Loaded.Feed.Path, adp)

		file, err := os.Open(adapterPath)
		if err != nil {
			config.Loaded.Logger.Errorw("error opening file", "adapter", adapterPath, "error", err)
		}

		fileStream, err := io.ReadAll(file)
		if err != nil {
			config.Loaded.Logger.Errorw("error reading file", "adapter", adapterPath, "error", err)
		}

		feed := NewAdapter()

		Import(fileStream, feed)

		adapters = append(adapters, feed)
	}

	return adapters
}

func Import(stream []byte, p *protoadapter.Adapter) {
	stream = []byte(os.ExpandEnv(string(stream)))

	if err := protojson.Unmarshal(stream, p); err != nil {
		config.Loaded.Logger.Fatal(err)
	}
}

func (dp *DataProvider) NewDataProviderService() (*grpc.Server, error) {
	var (
		opts []grpc.ServerOption
		s    = grpc.NewServer(opts...)
	)

	protoadapter.RegisterDataProviderServiceServer(s, dp)

	config.Loaded.Logger.Info("starting DP service on ", os.Getenv("HOST_IP"))

	return s, nil
}

func (dp *DataProvider) HandShake() (*grpc.ClientConn, error) {
	cfg, conn, error := newNodeServiceClient()
	client := cfg.(protonode.NodeServiceClient)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second) //Enough time to authenticate and add DP to known peers
	defer cancel()
	res, err := client.HandShake(ctx, &protonode.DPInfo{
		Name:          config.Loaded.Organization.Name,
		ListenAddress: dp.listenAddress,
		KOrgId:        config.Loaded.Organization.ID,
		Bootstraps:    dp.bootstraps,
		KnownPeers:    map[string]*protonode.DPInfo{},
	})

	if err != nil {
		log.Fatal("node.HandShake(_) failed: ", err)
	}

	if res.Status == true {
		config.Loaded.Logger.Info("node.HandShake(_) passed with response ", res.Status, " from ", config.Loaded.ServiceNode)
	}

	return conn, error
}

type OauthTokenSource struct {
	AuthKey string
}

func (s OauthTokenSource) Token() (token *oauth2.Token, err error) {
	return &oauth2.Token{
		AccessToken: s.AuthKey,
	}, nil
}

func newNodeServiceClient() (interface{}, *grpc.ClientConn, error) {
	oauthKey := os.Getenv("OAUTH_TOKEN")

	rpcCredentials := oauth.TokenSource{TokenSource: OauthTokenSource{
		AuthKey: oauthKey,
	}}

	nodeCertFilePath := path.Join(config.Loaded.RootWD, "certs/node", config.Loaded.SSL.Certificate)

	credentials, err := credentials.NewClientTLSFromFile(nodeCertFilePath, config.Loaded.Organization.Website)
	if err != nil {
		return nil, &grpc.ClientConn{}, fmt.Errorf("failed to load credentials: %v", err)
	}

	var opts []grpc.DialOption

	opts = []grpc.DialOption{
		grpc.WithPerRPCCredentials(rpcCredentials),
		grpc.WithTransportCredentials(credentials),
	}

	conn, err := grpc.Dial(config.Loaded.ServiceNode, opts...)
	if err != nil {
		return nil, &grpc.ClientConn{}, fmt.Errorf("did not connect: %v", err)
	}

	client := protonode.NewNodeServiceClient(conn)

	return client, conn, nil
}

func (dp *DataProvider) ListKnownPeers(ctx context.Context, null *protoadapter.Null) (*protoadapter.DPInfos, error) {

	dps := &protoadapter.DPInfos{}

	for _, dp := range dp.knownPeers {
		dps.List = append(dps.List, &protoadapter.DPInfo{
			ListenAddress: dp.ListenAddress,
			Name:          dp.Name,
			KOrgId:        dp.KOrgId,
		})
	}

	return dps, nil
}

func (dp *DataProvider) AddToKnownPeers(ctx context.Context, info *protoadapter.DPInfo) (*protoadapter.Null, error) {

	config.Loaded.Logger.Info("adding ", info.ListenAddress, " to known peers of bootstrap DP ", dp.listenAddress)

	dp.mu.Lock()
	defer dp.mu.Unlock()

	dpInfo := new(protonode.DPInfo)

	castBtwDPInfo(info, dpInfo)

	dp.knownPeers[info.ListenAddress] = dpInfo

	return new(protoadapter.Null), nil
}

func castBtwDPInfo(from, to interface{}) error {
	b, err := json.Marshal(from)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, to)
}
