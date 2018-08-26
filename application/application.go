package application

type Application struct {
	Config *Config
}

func Bootstrap(env string) *Application {
	return &Application{
		Config: LoadConfig(env),
	}
}
