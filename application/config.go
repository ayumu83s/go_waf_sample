package application

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Web   WebConfig
	Mysql MysqlConfig
}

type WebConfig struct {
	Port int
}

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
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
