package main

import (
	"log"
	"net"
	"os"
	"path"

	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klayoracle/klayoracle-monorepo/node/core"

	"github.com/klayoracle/klayoracle-monorepo/node/boot"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("cannot determine working directory: ", err)
	}

	boot.Boot(wd, path.Join(wd, "config.yaml"), path.Join(wd, ".env"))

	//Start Node service
	s, err := core.NewNodeServiceServer()
	if err != nil {
		config.Loaded.Logger.Fatal(err)
	}

	lis, err := net.Listen("tcp", os.Getenv("HOST_IP"))

	if err != nil {
		config.Loaded.Logger.Fatal("failed to listen: %v", err)
	}

	// @Todo Blocking action carried out, no code will run after this, use routine
	if err := s.Serve(lis); err != nil {
		s.Stop()
		config.Loaded.Logger.Fatal("failed to serve: %v", err)
	}
}
