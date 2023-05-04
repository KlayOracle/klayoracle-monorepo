package config

import (
	"log"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var Loaded Config

type Organization struct {
	ID      string `yaml:"k_org_id"`
	Name    string `yaml:"name"`
	Website string `yaml:"website"`
}

type Config struct {
	Env         string `yaml:"env"`
	RootWD      string
	Logger      *zap.SugaredLogger
	ServiceNode string `yaml:"service_node"`
	Feed        struct {
		Path     string   `yaml:"path"`
		Adapters []string `yaml:"adapters"`
	} `yaml:"feed"`
	Organization `yaml:"organization"`
	SSL          struct {
		Key         string `yaml:"key"`
		Certificate string `yaml:"certificate"`
	}
}

func Load(configFilePath string) {
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal("cannot open: ", configFile, "reason: ", err)
	}

	if err = yaml.Unmarshal(configFile, &Loaded); err != nil {
		log.Fatal("invalid config.yaml file", "reason: ", err)
	}

}
