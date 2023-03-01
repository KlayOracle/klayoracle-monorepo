package boot

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/joho/godotenv"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func Boot(wd, configPath, envPath string) {
	//Load env variables
	godotenv.Load(envPath)

	//Load yaml configs
	config.Load(configPath)
	config.Loaded.Logger = config.NewLogger(path.Join(wd, ".log"))
	config.Loaded.RootWD = wd
}

func HandShake() error {
	cfg, conn, error := newNodeServiceClient()
	client := cfg.(protonode.NodeServiceClient)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.HandShake(ctx, &protonode.DPInfo{
		Name:          config.Loaded.Organization.Name,
		ListenAddress: os.Getenv("HOST_IP"),
		KOrgId:        config.Loaded.Organization.ID,
	})

	if err != nil {
		log.Fatal("client.HandShake(_) failed: ", err)
	}

	fmt.Println(res.Status)

	return error
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
