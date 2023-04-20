package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	wd, _ := os.Getwd()

	buildDir := path.Join(wd, "..", "contracts", "build")
	targetDir := path.Join(wd, "..", "contracts", "oracle")
	abiGen := path.Join(wd, "..", "bin", "abigen")

	contractNames := []string{"KlayOracle", "KlayOracleInterface", "OracleProviderSample", "OracleConsumerSample"}

	for _, contractName := range contractNames {
		buildFilename := buildDir + "/" + contractName
		bindingFilename := targetDir + "/" + contractName + ".go"
		cmd := exec.Command(
			abiGen, fmt.Sprintf("--bin=%s.bin", buildFilename),
			fmt.Sprintf("--abi=%s.abi", buildFilename),
			fmt.Sprintf("--pkg=%s", contractName),
			fmt.Sprintf("--out=%s", bindingFilename),
		)

		err := cmd.Start()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Running... %s\n\n", cmd)

		err = cmd.Wait()

		if err != nil {
			log.Fatalln(err)
		}

		bindFile, err := os.ReadFile(bindingFilename)

		if err != nil {
			log.Fatalln("Cannot open: ", bindingFilename, err)
		}

		fileOutput := strings.Replace(string(bindFile), "package "+contractName, "package oracle", 1)

		err = os.WriteFile(bindingFilename, []byte(fileOutput), 0644)

		if err != nil {
			log.Fatalln("Cannot write to: ", bindingFilename, err)
		}
	}
}
