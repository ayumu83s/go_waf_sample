package application

import "database/sql"

type Application struct {
	Config *Config
	DB     *sql.DB
}

var app *Application

func Bootstrap(env string) *Application {
	config := LoadConfig(env)
	db := SetupDB(config)
	app = &Application{
		Config: config,
		DB:     db,
	}
	return app
}

func App() *Application {
	return app
}
