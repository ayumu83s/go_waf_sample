package application

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Web WebConfig
}

type WebConfig struct {
	Port int
}

func configPath(env string) string {
	return fmt.Sprintf("config/%s.toml", env)
}

func LoadConfig(env string) *Config {
	path := configPath(env)
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
