package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/adapter"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/boot"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot determine working directory: ", err)
	}

	//Load config, logger e.t.c
	boot.Boot(wd, path.Join(wd, "config.yaml"), path.Join(wd, ".env"))

	//Start Data Provider service
	s, err := adapter.NewDataProviderService()
	if err != nil {
		config.Loaded.Logger.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	connChan := make(chan *grpc.ClientConn)

	lis, err := net.Listen("tcp", os.Getenv("HOST_IP"))
	if err != nil {
		config.Loaded.Logger.Fatal("failed to listen: %v", err)
	}

	//Serve DP on port os.Getenv("HOST_IP")
	go func() {
		if err := s.Serve(lis); err != nil {
			cancel()
			config.Loaded.Logger.Fatal("failed to serve: %v", err)
		}
	}()

	//Create client and start handshake to Node Service
	go func() {
		conn, err := boot.HandShake()
		connChan <- conn

		if err != nil {
			cancel()
			config.Loaded.Logger.Fatal("DP Handshake: ", err)
		}
	}()

	for {
		select {
		case conn := <-connChan:
			config.Loaded.Logger.Info("Closing client connection to ", conn.Target())
			conn.Close()
		case <-ctx.Done(): //If DP Server crashes or Handshake fails
			s.Stop() //Don't take chances with resources and be sure DP Service closes
			fmt.Println("Data provider operation... exited")
			return
		}
	}
}
