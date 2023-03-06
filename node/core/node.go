package core

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"

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
	errAddingToKnownPeer          = status.Errorf(codes.Unknown, "cannot add to known dp known peers")
)

type Node struct {
	protonode.UnimplementedNodeServiceServer
	Sever         *grpc.Server
	dataProviders map[string]*protonode.DPInfo
	mu            sync.Mutex
}

func (n *Node) HandShake(ctx context.Context, provider *protonode.DPInfo) (*protonode.HandShakeStatus, error) {

	//@Todo confirm organization is whitelisted in DB
	//Node runner uses cmd to add and remove from supported organization
	if isDPWhitelist(provider) != true {
		return nil, errDataProviderNotWhitelisted
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	config.Loaded.Logger.Info("Registering provider ", provider.ListenAddress, " with Node ", os.Getenv("HOST_IP"))

	n.dataProviders[provider.ListenAddress] = provider

	config.Loaded.Logger.Info("Adding provider to known peers ", provider.ListenAddress)

	err := addPeer(provider)
	if err != nil {
		config.Loaded.Logger.Fatal("error adding to known peer: ", err)
		return nil, errAddingToKnownPeer
	}

	//Check DP Peers to find out Org is not yet registered
	return &protonode.HandShakeStatus{Status: true, ErrorMsg: ""}, nil
}

func addPeer(p *protonode.DPInfo) error {

	//Send a grpc request to all bootstrap nodes for Known peers
	//Identify the bootstrap dp node with the longest known peers chain
	// Add new dp Info to its peers and loop through the all bootstrap nodes to update the list of known peers

	//Gather known peers from all bootstrap dp node
	//Identify the longest peer chain
	var (
		longestChain   = 0
		longestChainDP = ""
		peerList       = map[string]*protoadapter.DPInfos{}
		peers          = &protoadapter.DPInfos{}
		err            error
	)

	for listenAddr, _ := range p.Bootstraps {
		config.Loaded.Logger.Info("Fetching know peers for bootstrap dp: ", listenAddr)
		peers, err = getPeerList(listenAddr)
		if err != nil {
			return err
		}

		peerList[listenAddr] = peers

		if len(peers.List) > longestChain {
			longestChain = len(peers.List)
			longestChainDP = listenAddr
		}

		config.Loaded.Logger.Info("Found peers: ")
		config.Loaded.Logger.Info(peers)
	}

	config.Loaded.Logger.Infow("dp bootstrap with the longest peer", longestChainDP, longestChain)

	//list := append(peerList[longestChainDP].List, &protoadapter.DPInfo{
	//	ListenAddress: p.ListenAddress,
	//	KOrgId:        p.KOrgId,
	//})

	//bootstraps := bootstrap.Nodes()

	//for i, _ := range list {
	//dp := peerList[longestChainDP].List[i]
	//
	////Bootstrap node can't add self as peer
	//if slices.Contains(bootstraps, p.ListenAddress) {
	//	continue
	//}

	//client, err := newDPClient(dp.ListenAddress)
	//if err != nil {
	//	return fmt.Errorf("failed connecting to client: %v", err)
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Enough time to authenticate and add DP to known peers
	//defer cancel()

	//res, err := client.Add(ctx, &protoadapter.Null{})
	//
	//if err != nil {
	//	return fmt.Errorf("client.ListKnowPeers(_) failed: %v", err)
	//}
	//}

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
		return nil, fmt.Errorf("client.ListKnowPeers(_) failed: %v", err)
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

func NewNodeServiceServer() (*grpc.Server, error) {
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

	protonode.RegisterNodeServiceServer(s, &Node{
		dataProviders: make(map[string]*protonode.DPInfo),
		Sever:         s,
	})

	config.Loaded.Logger.Info("Starting Node Service on ", os.Getenv("HOST_IP"))

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
