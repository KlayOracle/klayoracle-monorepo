package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/klayoracle/klayoracle-monorepo/node/core"

	"github.com/klayoracle/klayoracle-monorepo/node/boot"
	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	wd, _ := os.Getwd()
	boot.Boot(path.Join(wd, "../../node"), path.Join(wd, "../../node", "config.yaml"), path.Join(wd, "../../node", ".env"))

	flag.Parse()
	adapterList := flag.Args()

	cfg := config.Loaded
	logger := cfg.Logger

	if len(adapterList) < 1 {
		logger.Infow("no args passed", "example command: ", "make adapter-dry-run ADAPTERS=KLAY_USD.json WEMIX_USD.json")
	}

	for _, s := range adapterList {
		adapterPath := path.Join(cfg.RootWD, "..", "data-provider", "feeds", s)

		logger.Infow("loading adapter", "path", adapterPath)

		file, err := os.Open(adapterPath)

		if err != nil {
			logger.Errorw("error opening file", "adapter", adapterPath, "error", err)
		}

		fileStream, err := io.ReadAll(file)
		if err != nil {
			logger.Errorw("error reading file", "adapter", adapterPath, "error", err)
		}

		newAdapter := protonode.Adapter{}
		stream := []byte(os.ExpandEnv(string(fileStream)))
		protojson.Unmarshal(stream, &newAdapter)

		result, err := core.DryRun(newAdapter)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n")

		logger.Infow("ok!", "feed", newAdapter.Name, "output", result)

		fmt.Printf("\n\n")
	}
}
