package main

import (
	"fmt"
	"log"

	"github.com/ayumu83s/di_sample/application"
	"github.com/ayumu83s/di_sample/controller"
	"github.com/ayumu83s/di_sample/registory"
)

func main() {
	app := application.Bootstrap("dev")
	fmt.Println("port: ", app.Config.Web.Port)

	var name string
	if err := app.DB.QueryRow("SELECT name FROM user WHERE id = ?", 1).Scan(&name); err != nil {
		log.Fatal(err)
	}
	fmt.Println("name: ", name)

	defer app.DB.Close() // これをCloseするのはどこが正解なのか・・

	repo := registory.NewUserRepository()
	controller := controller.NewUser(repo)
	controller.Get(1)
}
