package core

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

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

	err := addPeer(provider)

	if err != nil {
		return nil, nil
	}

	//Check DP Peers to find out Org is not yet registered
	return &protonode.HandShakeStatus{Status: true, ErrorMsg: ""}, nil
}

func addPeer(p *protonode.DPInfo) error {
	//for listenAddr, _ := range p. {
	//
	//}

	return fmt.Errorf("")
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
