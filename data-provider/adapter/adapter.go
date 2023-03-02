package adapter

import (
	"fmt"
	"os"
	"sync"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/bootstrap"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/protoadapter"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type DataProvider struct {
	protoadapter.UnimplementedDataProviderServiceServer
	bootstraps map[string]*protoadapter.DPInfo // A list of 3 default bootstrap dp nodes
	knownPeers map[string]*protoadapter.DPInfo // A list of dp nodes that are not bootstrap nodes, and only known to bootstrap nodes
	mu         sync.Mutex
}

func NewAdapter() *protoadapter.Adapter {
	return &protoadapter.Adapter{}
}

func Import(stream []byte, p *protoadapter.Adapter) {
	stream = []byte(os.ExpandEnv(string(stream)))

	if err := protojson.Unmarshal(stream, p); err != nil {
		config.Loaded.Logger.Fatal(err)
	}
}

func NewDataProviderService() (*grpc.Server, error) {
	var (
		opts []grpc.ServerOption
		s    = grpc.NewServer(opts...)
	)

	bootstraps := make(map[string]*protoadapter.DPInfo)
	bootstrapNodes := bootstrap.Nodes()

	for i, dp := range bootstrapNodes {
		bootstraps[dp] = &protoadapter.DPInfo{
			Name:          fmt.Sprintf("Bootstrap-00%d", i),
			ListenAddress: dp,
			KOrgId:        fmt.Sprintf("bootstrap-00%d", i),
		}
	}

	protoadapter.RegisterDataProviderServiceServer(s, &DataProvider{
		knownPeers: make(map[string]*protoadapter.DPInfo),
		bootstraps: bootstraps,
	})

	config.Loaded.Logger.Info("Starting DP Service on ", os.Getenv("HOST_IP"))

	return s, nil
}
