package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"

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

	config.Loaded.Logger.Info("Working directory: ", wd)

	boot.Boot(wd, path.Join(wd, "config.yaml"), path.Join(wd, ".env"))

	//Start Node service
	n := &core.Node{}
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
		if err := s.Serve(lis); err != nil {
			shutServer <- s
		}
	}()

	for {
		select {
		case job := <-n.Jobs:
			fmt.Printf("Adapter Job: %v", job)
		case <-shutServer: //If DP Server crashes or Handshake fails
			s.Stop()
			config.Loaded.Logger.Fatal("failed to serve: %v", err)
			return
		}
	}
}
