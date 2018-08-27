package main

import (
	"fmt"

	"github.com/ayumu83s/go_waf_sample/application"
	"github.com/ayumu83s/go_waf_sample/controller"
	"github.com/ayumu83s/go_waf_sample/registory"
)

func main() {
	app := application.Bootstrap("dev")
	fmt.Println("port: ", app.Config.Web.Port)

	defer app.DB.Close() // これをCloseするのはどこが正解なのか・・

	repoArgs := &registory.RepositoryInput{
		DB: app.DB,
	}
	repo := registory.NewUserRepository(repoArgs)
	controller := controller.NewUser(repo)
	controller.Get(1)
	controller.Get(2)
	controller.Get(99)
}
