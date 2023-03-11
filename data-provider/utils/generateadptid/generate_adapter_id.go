package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"io"
	"log"
	"os"
	"path"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/adapter"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/boot"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	wd, _ := os.Getwd()
	boot.Boot(path.Join(wd, ".."), path.Join(wd, "..", "config.yaml"), path.Join(wd, "..", ".env"))

	flag.Parse()
	adapterList := flag.Args()

	cfg := config.Loaded
	logger := cfg.Logger

	if len(adapterList) < 1 {
		logger.Infow("no args passed", "example command: ", "ADAPTERS=KLAY_USD.json WEMIX_USD.json make adapter-id-gen")
	}

	for _, s := range adapterList {
		adapterPath := path.Join(cfg.RootWD, cfg.Feed.Path, s)
		logger.Infow("loading adapter", "path", adapterPath)

		file, err := os.Open(adapterPath)

		if err != nil {
			logger.Errorw("error opening file", "adapter", adapterPath, "error", err)
		}

		fileStream, err := io.ReadAll(file)
		if err != nil {
			logger.Errorw("error reading file", "adapter", adapterPath, "error", err)
		}

		feed := adapter.NewAdapter()
		adapter.Import(fileStream, feed)

		feed.AdapterId = generateRandomHex()

		updatedStream, err := protojson.Marshal(feed)
		if err != nil {
			logger.Errorw("error marshaling", "adapter", adapterPath, "error", err)
		}

		logger.Infow("updated adapter", "file", s, "adapter_id", feed.AdapterId)

		err = os.WriteFile(adapterPath, updatedStream, 0644)
		if err != nil {
			logger.Errorw("error updating adapter id", "adapter", adapterPath, "error", err)
		}

		err = file.Close()
		if err != nil {
			logger.Errorw("error closing adapter file", "adapter", adapterPath, "error", err)
		}
	}
}

func generateRandomHex() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return "0x" + hex.EncodeToString(bytes)
}
