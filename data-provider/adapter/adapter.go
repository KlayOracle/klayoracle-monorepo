package adapter

import (
	"os"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/protoadapter"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"google.golang.org/protobuf/encoding/protojson"
)

type DataProvider struct {
	protoadapter.UnimplementedDataProviderServer
	NodeServiceClient *protonode.NodeServiceClient
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
