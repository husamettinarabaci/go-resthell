package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type RestConfig struct {
	Debug   bool `yaml:"debug"`
	Restapi struct {
		Port string `yaml:"port"`
	} `yaml:"restapi"`
}

var RestConfigInstance *RestConfig

func GetRestConfigInstance() *RestConfig {
	if RestConfigInstance == nil {
		RestConfigInstance = &RestConfig{}
		RestConfigInstance.ReadConfig()
	}
	return RestConfigInstance
}

var RestConfigPath string = "config/rest.yml"
var RestConfigLocalPath string = "config/rest_local.yml"

func (c *RestConfig) ReadConfig() {

	configPath := RestConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = RestConfigLocalPath
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

func (c *RestConfig) IsDebug() bool {
	return GetRestConfigInstance().Debug
}
