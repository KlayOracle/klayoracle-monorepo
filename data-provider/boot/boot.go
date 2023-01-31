package boot

import (
	"log"
	"os"
	"path"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"
)

func Boot() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot determine working directory: ", err)
	}

	config.Load(path.Join(wd, "config.yaml"))
	config.Loaded.Logger = config.NewLogger(path.Join(wd, ".log"))
	config.Loaded.RootWD = wd
}
