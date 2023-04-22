package boot

import (
	"path"

	"github.com/klayoracle/klayoracle-monorepo/node/core"

	"github.com/klayoracle/klayoracle-monorepo/node/storage"

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

	core.NewClient()

	storage.CreateDBConn()
}
