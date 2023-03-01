package boot

import (
	"net"
	"os"
	"path"

	"github.com/klayoracle/klayoracle-monorepo/node/core"

	"github.com/joho/godotenv"

	"github.com/klayoracle/klayoracle-monorepo/node/config"
)

func Boot(wd, configPath, envPath string) {
	//Load env variables
	godotenv.Load(envPath)

	//Load yaml configs
	config.Load(configPath)
	config.Loaded.Logger = config.NewLogger(path.Join(wd, ".log"))
	config.Loaded.RootWD = wd

	//Start Node service
	s, err := core.NewNodeServiceServer()
	if err != nil {
		config.Loaded.Logger.Fatal(err)
	}

	lis, err := net.Listen("tcp", os.Getenv("HOST_IP"))

	if err != nil {
		config.Loaded.Logger.Fatal("failed to listen: %v", err)
	}

	node := core.Node{}
	node.Sever = s

	// @Todo Blocking action carried out, no code will run after this, use routine
	if err := s.Serve(lis); err != nil {
		config.Loaded.Logger.Fatal("failed to serve: %v", err)
	}
}
