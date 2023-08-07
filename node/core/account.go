package core

import (
	"context"
	"os"

	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klaytn/klaytn/client"
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
