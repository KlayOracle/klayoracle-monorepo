package config

import (
	"encoding/json"
	"go.uber.org/zap"
)

func NewLogger(logFilePath string) *zap.SugaredLogger {

	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "console",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	cfg.OutputPaths = append(cfg.OutputPaths, logFilePath)

	return zap.Must(cfg.Build()).Sugar()
}
