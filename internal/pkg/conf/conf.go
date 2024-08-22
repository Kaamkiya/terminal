package conf

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	BannerPath string `yaml:"banner_path"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func LoadConfig(configFilePath string) Config {
	config := Config{}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return config
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		return config
	}

	return config
}
