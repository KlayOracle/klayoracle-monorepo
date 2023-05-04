package main

import (
	"log"
	"net"
	"os"
	"path"

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
	core.NewClient()
	defer func() {
		core.KlaytnClient.Close()
	}()

	config.Loaded.Logger.Info("Working directory: ", wd)

	//Start Node service
	n := &core.Node{}
	n.Jobs = make(map[string]*protonode.Adapter) //Fix nil map issue
	s, err := core.NewNodeServiceServer(n)
	if err != nil {
		config.Loaded.Logger.Fatal(err)
	}

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
