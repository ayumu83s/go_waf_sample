package main

import (
	"fmt"

	"github.com/ayumu83s/di_sample/application"
	"github.com/ayumu83s/di_sample/controller"
	"github.com/ayumu83s/di_sample/registory"
)

func main() {
	config := application.LoadConfig("config/dev.toml")
	fmt.Println("port: ", config.Web.Port)
	repo := registory.NewUserRepository()
	controller := controller.NewUser(repo)
	controller.Get(1)
}
