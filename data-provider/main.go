package main

import (
	"github.com/klayoracle/klayoracle-monorepo/data-provider/boot"
	"log"
	"os"
	"path"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot determine working directory: ", err)
	}

	boot.Boot(wd, path.Join(wd, "config.yaml"), path.Join(wd, ".env"))
}
