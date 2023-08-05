package core

import (
	"context"
	"os"

	client "github.com/ethereum/go-ethereum/ethclient"
	"github.com/klayoracle/klayoracle-monorepo/node/config"
)

var (
	KlaytnClient    *client.Client
	KlaytnClientCtx context.Context
	KlaytnChainId   = os.Getenv("CHAIN_ID")
)

func NewWssClient() {
	client, err := client.Dial(os.Getenv("NODE_URL_WSS"))
	if err != nil {
		config.Loaded.Logger.Fatalw("wss connection to blockchain client failed", "error", err)
	}

	KlaytnClient = client
	KlaytnClientCtx = context.Background()
}

func NewHttpClient() (*client.Client, error) {
	return client.Dial(os.Getenv("NODE_URL_HTTPS"))
}
