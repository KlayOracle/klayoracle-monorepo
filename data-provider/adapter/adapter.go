package adapter

import (
	"os"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/protoadapter"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewAdapter() *protoadapter.Adapter {
	return &protoadapter.Adapter{}
}

func Import(stream []byte, p *protoadapter.Adapter) {
	stream = []byte(os.ExpandEnv(string(stream)))

	if err := protojson.Unmarshal(stream, p); err != nil {
		config.Loaded.Logger.Fatal(err)
	}
}
