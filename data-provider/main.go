package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/klayoracle/klayoracle-monorepo/node/protonode"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/adapter"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/boot"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("cannot determine working directory: ", err)
	}

	if os.Getenv("WORK_DIR") != "" {
		wd = os.Getenv("WORK_DIR")
	}

	//Load config, logger e.t.c
	boot.Boot(wd, path.Join(wd, "config.yaml"), path.Join(wd, ".env"))

	config.Loaded.Logger.Info("Working directory: ", wd)

	dp := adapter.NewDataProvider()

	//Start Data Provider service
	s, err := dp.NewDataProviderService()
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
		conn, err := dp.HandShake()
		connChan <- conn

		if err != nil {
			cancel()
			config.Loaded.Logger.Fatal("DP Handshake: ", err)
		}

		adapters := adapter.ListAdapters()

		config.Loaded.Logger.Infow("send feeds to node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "total", len(adapters))

		for _, adapterCfg := range adapters {
			adapterCfg := adapterCfg

			go func() {
				ticker := time.NewTicker(time.Duration(adapterCfg.Frequency))

				for t := range ticker.C {
					config.Loaded.Logger.Infow("sending adapter request to service node", "timer", t, "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "adapter", adapterCfg.AdapterId, "name", adapterCfg.Name)

					func() {
						ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
						defer cancel()

						ctx = context.WithValue(ctx, "provider", os.Getenv("HOST_IP"))
						client := protonode.NewNodeServiceClient(conn)

						stream, err := client.QueueJob(ctx)
						if err != nil {
							config.Loaded.Logger.Fatalw("error sending adapter request to service node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "error", err)
						}

						nodeAdapter := new(protonode.Adapter)

						adapter.CastBtwDPInfo(adapterCfg, nodeAdapter)

						fmt.Println(nodeAdapter)
						fmt.Println("Stream: ", stream)
						//if err = stream.Send(nodeAdapter); err != nil {
						//	config.Loaded.Logger.Warnw("error sending adapter request to service node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "error", err)
						//}
						//
						//oracleRequest, err := stream.Recv()
						//if err != nil {
						//	config.Loaded.Logger.Warnw("error receiving oracle request response from service node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "error", err)
						//}
						//
						//if err = stream.CloseSend(); err != nil {
						//	config.Loaded.Logger.Warnw("error receiving oracle request response from service node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "error", err)
						//}
						//
						//fmt.Println("Oracle Request: ", oracleRequest)

					}()
				}
			}()
		}
	}()

	for {
		select {
		case conn := <-connChan:
			config.Loaded.Logger.Info("closing client connection to ", conn.Target())
			//conn.Close()
		case <-ctx.Done(): //If DP Server crashes or Handshake fails
			s.Stop() //Don't take chances with resources and be sure DP Service closes
			fmt.Println("data provider operation... exited")
			return
		}
	}
}
