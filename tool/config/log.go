package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type LogConfig struct {
	Debug  bool `yaml:"debug"`
	Logger struct {
		Console bool `yaml:"console"`
	} `yaml:"logger"`
}

var LogConfigInstance *LogConfig

func GetLogConfigInstance() *LogConfig {
	if LogConfigInstance == nil {
		LogConfigInstance = &LogConfig{}
		LogConfigInstance.ReadConfig()
	}
	return LogConfigInstance
}

var LogConfigPath string = "config/log.yml"
var LogConfigLocalPath string = "config/log_local.yml"

func (c *LogConfig) ReadConfig() {

	configPath := LogConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = LogConfigLocalPath
		}
	}
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

func (c *LogConfig) IsDebug() bool {
	return GetLogConfigInstance().Debug
}
