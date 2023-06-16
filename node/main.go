package main

import (
	"github.com/klaytn/klaytn/common"
	"log"
	"math/big"
	"net"
	"os"
	"path"
	"time"

	"github.com/klayoracle/klayoracle-monorepo/node/protonode"

	"google.golang.org/grpc"

	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klayoracle/klayoracle-monorepo/node/core"

	"github.com/klayoracle/klayoracle-monorepo/node/boot"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("cannot determine working directory: ", err)
	}

	if os.Getenv("WORK_DIR") != "" {
		wd = os.Getenv("WORK_DIR")
	}

	boot.Boot(wd, path.Join(wd, "config.yaml"), path.Join(wd, ".env"))

	//@Todo refactor
	core.NewWssClient()
	defer func() {
		config.Loaded.Logger.Warnw("closing blockchain client connection")
		core.KlaytnClient.Close()
	}()

	config.Loaded.Logger.Info("Working directory: ", wd, "\n")

	//Start Node service
	n := &core.Node{}
	n.Jobs = make(map[string]*protonode.Adapter) //Fix nil map issue
	s, err := core.NewNodeServiceServer(n)
	if err != nil {
		config.Loaded.Logger.Fatal(err)
	}

	//Set round queue default
	n.RoundQueue = new([]core.RoundQueue)
	nodeAddress := common.HexToAddress(os.Getenv("PUBLIC_ADDRESS"))
	nonce, err := core.KlaytnClient.NonceAt(core.KlaytnClientCtx, nodeAddress, nil)
	if err != nil {
		config.Loaded.Logger.Fatal("cannot determine nonce for node address: %v", err)
	}

	n.LastNonce = big.NewInt(int64(nonce))

	//Start watching queue every 5 secs
	go func() {
		for now := range time.Tick(time.Second * 5) {
			config.Loaded.Logger.Infow("checking round queue for available answer", "size", len(*n.RoundQueue), "time", now.Local())
			n.WatchRoundQueue()
		}
	}()

	lis, err := net.Listen("tcp", os.Getenv("HOST_IP"))
	if err != nil {
		config.Loaded.Logger.Fatal("failed to listen: %v", err)
	}

	shutServer := make(chan *grpc.Server, 1)

	go func() {
		if err = s.Serve(lis); err != nil {
			shutServer <- s
		}
	}()

	n.Peer()

	for range shutServer {
		s.Stop()
		config.Loaded.Logger.Fatal("failed to serve: %v", err)
		return
	}
}
