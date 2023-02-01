package boot

import (
	"path"

	"github.com/joho/godotenv"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
)

func Boot(wd, configPath, envPath string) {
	godotenv.Load(envPath)

	config.Load(configPath)
	config.Loaded.Logger = config.NewLogger(path.Join(wd, ".log"))
	config.Loaded.RootWD = wd
}
