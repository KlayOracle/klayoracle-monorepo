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
	KlaytnChainId   = 1001
)

func NewClient() {
	client, err := client.Dial(os.Getenv("NODE_URL"))
	if err != nil {
		config.Loaded.Logger.Fatalw("connection to klaytn client failed", "error", err)
	}

	KlaytnClient = client
	KlaytnClientCtx = context.Background()
}
