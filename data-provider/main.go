package main

import (
	"log"
	"os"
	"path"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/config"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/boot"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot determine working directory: ", err)
	}

	boot.Boot(wd, path.Join(wd, "config.yaml"), path.Join(wd, ".env"))

	//Create client and start handshake to Node
	err = boot.HandShake()
	if err != nil {
		config.Loaded.Logger.Fatal(err)
	}
}
