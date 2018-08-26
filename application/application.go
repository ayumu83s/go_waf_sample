package application

import "database/sql"

type Application struct {
	Config *Config
	DB     *sql.DB
}

func Bootstrap(env string) *Application {
	config := LoadConfig(env)
	db := SetupDB(config)
	return &Application{
		Config: config,
		DB:     db,
	}
}
