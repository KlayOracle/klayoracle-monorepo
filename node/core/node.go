package core

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/klayoracle/klayoracle-monorepo/node/storage"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"golang.org/x/exp/slices"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/bootstrap"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/protoadapter"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata            = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken               = status.Errorf(codes.Unauthenticated, "invalid token")
	errDataProviderNotWhitelisted = status.Errorf(codes.Unknown, "you need to be whitelisted")
	errAddingToKnownPeer          = status.Errorf(codes.Unknown, "cannot add to known DP known peers")
)

type Node struct {
	protonode.UnimplementedNodeServiceServer
	Sever         *grpc.Server
	dataProviders map[string]*protonode.DPInfo
	mu            sync.Mutex
	Jobs          []*protonode.Adapter //Using multiple routines to R&W to channel is safe without mutex.
}

func (n *Node) HandShake(ctx context.Context, provider *protonode.DPInfo) (*protonode.HandShakeStatus, error) {

	//@Todo confirm organization is whitelisted in DB
	//Node runner uses cmd to add and remove from supported organization
	if isDPWhitelist(provider) != true {
		return nil, errDataProviderNotWhitelisted
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	config.Loaded.Logger.Infow("attempting to add provider to all bootstrap dp known peers ", "listen address", provider.ListenAddress)

	if slices.Contains(bootstrap.Nodes(), provider.ListenAddress) {
		config.Loaded.Logger.Warnw("ignoring adding to known peer", "reason", "bootstrap dp", "listen address", provider.ListenAddress)
	} else {
		err := addPeer(provider)
		if err != nil {
			config.Loaded.Logger.Fatal("error adding to known peer: ", err)
			return nil, errAddingToKnownPeer
		}
	}

	config.Loaded.Logger.Infow("registering provider with service node ", "provider", provider.ListenAddress, "service node", os.Getenv("HOST_IP"))

	n.dataProviders[provider.ListenAddress] = provider

	//Check DP Peers to find out Org is not yet registered
	return &protonode.HandShakeStatus{Status: true, ErrorMsg: ""}, nil
}

func (n *Node) QueueJob(ctx context.Context, adapter *protonode.Adapter) (*protonode.RequestStatus, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	provider := md["provider"][0]

	_, ok := n.dataProviders[provider]

	if ok {
		config.Loaded.Logger.Infow("new queue request", "data provider", provider, "request", adapter)
		n.mu.Lock()
		defer n.mu.Unlock()
		n.Jobs = append(n.Jobs, adapter)

		storage.StoreJob(provider, adapter)

		config.Loaded.Logger.Info("Job in queue: ", len(n.Jobs))
	} else {
		config.Loaded.Logger.Warnw("adapter has not registered an handshake", "service node",
			os.Getenv("HOST_IP"), "data provider", provider)

		return &protonode.RequestStatus{
			Status: 1,
		}, fmt.Errorf("adapter has not registered an handshake")
	}

	return &protonode.RequestStatus{
		Status: 0,
	}, nil
}

func addPeer(p *protonode.DPInfo) error {

	//Send a grpc request to all bootstrap nodes for Known peers
	//Identify the bootstrap dp node with the longest known peers chain
	//Add new dp Info to its peers and loop through the all bootstrap nodes to update the list of known peers

	//Gather known peers from all bootstrap dp node
	//Identify the longest peer chain
	var (
		longestChain   = 0                                  //Count of the longest chain
		longestChainDP = ""                                 //Listen Ip of the longest DP's peer-chain
		peerList       = map[string]*protoadapter.DPInfos{} //Mapping of bootstrap dp listen address to their dp peers
		peers          = &protoadapter.DPInfos{}            //
		err            error
	)

	bootstraps := bootstrap.Nodes()

	//@Todo If a bootstrap node does not respond, it should not stop a new peer from responding
	for _, listenAddr := range bootstraps {
		config.Loaded.Logger.Info("fetching known peers for bootstrap DP: ", listenAddr)
		peers, err = getPeerList(listenAddr)
		if err != nil {
			config.Loaded.Logger.Warn("bootstrap DP is unresponsive on ", listenAddr)

			//return err
		} else {

			peerList[listenAddr] = peers

			if len(peers.List) >= longestChain {
				longestChain = len(peers.List)
				longestChainDP = listenAddr
			}

			config.Loaded.Logger.Info("peers found for ", listenAddr)
			if len(peers.List) > 0 {
				config.Loaded.Logger.Info(peers)
			} else {
				config.Loaded.Logger.Warn("No known peers discovered")
			}
		}
	}

	config.Loaded.Logger.Infow("DP bootstrap with the longest peer", longestChainDP, longestChain)

	dp := &protoadapter.DPInfo{}

	err = castBtwDPInfo(p, dp)

	if err != nil {
		config.Loaded.Logger.Fatal("conversion between DPInfo gone wrong: ", err)
	}

	list := append(peerList[longestChainDP].List, dp)

	for _, bt := range bootstraps {

		//Due to defer behaviour in loop, we do this instead
		err := func() error {

			//Bootstrap node can't add self as peer
			client, err := newDPClient(bt)
			if err != nil {
				return fmt.Errorf("failed connecting to client: %v", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) //Enough time to authenticate and add DP to known peers
			defer cancel()

			//Add all DP services in list as this bootstrap DP know peer
			for _, newDp := range list {
				_, err = client.AddToKnownPeers(ctx, newDp)

				if err != nil {
					config.Loaded.Logger.Warnw("dp.AddToKnownPeers(_) failed", "reason", "unresponsive bootstrap dp", "skipping", bt)
				} else {
					config.Loaded.Logger.Infow("Added dp to bootstrap dp known peer ", "boostrap dp", bt, "added peer", newDp.ListenAddress)
				}
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

func getPeerList(listenAddr string) (*protoadapter.DPInfos, error) {

	client, err := newDPClient(listenAddr)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Enough time to authenticate and add DP to known peers
	defer cancel()

	res, err := client.ListKnownPeers(ctx, &protoadapter.Null{})

	if err != nil {
		return nil, fmt.Errorf("dp.ListKnowPeers(_) failed: %v", err)
	}

	return res, nil
}

func newDPClient(listenAddr string) (protoadapter.DataProviderServiceClient, error) {
	conn, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("dial failed: %v", err)
	}

	return protoadapter.NewDataProviderServiceClient(conn), nil
}

func NewNodeServiceServer(n *Node) (*grpc.Server, error) {
	certFilePath := path.Join(config.Loaded.RootWD, "certs/host", config.Loaded.SSL.Certificate)
	keyFilePath := path.Join(config.Loaded.RootWD, "certs/host", config.Loaded.SSL.Key)

	cert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return new(grpc.Server), fmt.Errorf("failed to load credentials: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(opts...)

	n.dataProviders = make(map[string]*protonode.DPInfo)
	n.Sever = s

	protonode.RegisterNodeServiceServer(s, n)

	config.Loaded.Logger.Info("starting Node service on ", os.Getenv("HOST_IP"))

	return s, nil
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	return token == os.Getenv("OAUTH_TOKEN")
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

func isDPWhitelist(dpi *protonode.DPInfo) bool {
	return true
}

func castBtwDPInfo(from, to interface{}) error {
	b, err := json.Marshal(from)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, to)
}
