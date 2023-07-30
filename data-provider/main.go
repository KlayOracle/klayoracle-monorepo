package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/klayoracle/klayoracle-monorepo/node/protonode"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/adapter"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/boot"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"golang.org/x/net/context"
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
		hConn, err := dp.HandShake() //handshake connection
		defer func() {
			err = hConn.Close()
			if err != nil {
				config.Loaded.Logger.Warnw("cannot close connection", "error", err)
			}
		}()

		if err != nil {
			cancel()
			config.Loaded.Logger.Fatal("DP Handshake: ", err)
		}

		adapters := adapter.ListAdapters(true)

		config.Loaded.Logger.Infow("send feeds to node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "total", len(adapters))

		for _, adapterCfg := range adapters {
			adapterCfg := adapterCfg

			go func() {
				ticker := time.NewTicker(time.Duration(adapterCfg.Frequency))

				for t := range ticker.C {
					config.Loaded.Logger.Infow("sending adapter request to service node", "timer", t, "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "adapter", adapterCfg.AdapterId, "name", adapterCfg.Name)

					func() {
						cfg, conn, err := adapter.NewNodeServiceClient()
						defer func() {
							err = conn.Close()
							if err != nil {
								config.Loaded.Logger.Warnw("cannot close connection", "error", err)
							}
						}()

						if err != nil {
							config.Loaded.Logger.Warnw("error sending adapter request to service node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "error", err)
						}

						client := cfg.(protonode.NodeServiceClient)

						ctx = context.Background()
						md := metadata.Pairs("provider", os.Getenv("HOST_IP"))
						ctx = metadata.NewOutgoingContext(ctx, md)

						nodeAdapter := new(protonode.Adapter)
						err = adapter.CastBtwDPInfo(adapterCfg, nodeAdapter)
						if err != nil {
							config.Loaded.Logger.Warnw("error sending adapter request to service node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "error", err)
						} else {

							status, err := client.QueueJob(ctx, nodeAdapter)

							if err != nil || status.Status == 1 {
								config.Loaded.Logger.Warnw("error sending adapter request to service node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "error", err)

								//re-attempt HandShake with node again until reconnect
								if hConn != nil {
									_ = hConn.Close()
								}

								config.Loaded.Logger.Infow("re attempting handshake to node after 60secs", "node", config.Loaded.ServiceNode)

								hConn, err = dp.HandShake() //handshake connection

								if err != nil {
									config.Loaded.Logger.Infow("could not re-established connection with node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode, "error", err)
								} else {
									config.Loaded.Logger.Infow("connection re-established with node", "data provider", os.Getenv("HOST_IP"), "node", config.Loaded.ServiceNode)
								}

							}
						}

					}()
				}
			}()
		}
	}()

	for range ctx.Done() {
		s.Stop() //Don't take chances with resources and be sure DP Service closes
		fmt.Println("data provider operation... exited")
		return
	}
}
