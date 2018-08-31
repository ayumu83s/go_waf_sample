package application

import (
	"database/sql"
	"github.com/garyburd/redigo/redis"
)

type Application struct {
	Config *Config
	DB     *sql.DB
	Redis  redis.Conn
}

var app *Application

func Bootstrap(env string) *Application {
	config := LoadConfig(env)
	db := SetupDB(config)
	redis := SetupRedis(config)

	app = &Application{
		Config: config,
		DB:     db,
		Redis: redis,
	}
	return app
}

func App() *Application {
	return app
}
