package application

import "github.com/BurntSushi/toml"

type Config struct {
	Web WebConfig
}

type WebConfig struct {
	Port int
}

func LoadConfig(path string) *Config {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
