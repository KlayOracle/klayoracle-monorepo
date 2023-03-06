package boot

import (
	"github.com/joho/godotenv"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
	"path"
)

func Boot(wd, configPath, envPath string) {
	//Load env variables
	godotenv.Load(envPath)

	//Load yaml configs
	config.Load(configPath)
	config.Loaded.Logger = config.NewLogger(path.Join(wd, ".log"))
	config.Loaded.RootWD = wd
}
