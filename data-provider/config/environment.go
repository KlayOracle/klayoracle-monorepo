package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Loaded Config

type Config struct {
	Env         string `yaml:"env"`
	RootWD      string
	Logger      *zap.SugaredLogger
	ServiceNode string `yaml:"service_node"`
	Feed        struct {
		Path     string   `yaml:"path"`
		Adapters []string `yaml:"adapters"`
	} `yaml:"feed"`
	Organization struct {
		ID      string `yaml:"k_org_id"`
		Name    string `yaml:"name"`
		Website string `yaml:"website"`
	} `yaml:"organization"`
	SSL struct {
		Key         string `yaml:"key"`
		Certificate string `yaml:"certificate"`
	}
}

func Load(configFilePath string) {
	configFile, err := os.ReadFile(configFilePath)

	err = yaml.Unmarshal(configFile, &Loaded)
	if err != nil {
		log.Fatal("Invalid config.yaml file", "reason", err)
	}

}
