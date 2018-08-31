package main

import (
	"fmt"

	"github.com/ayumu83s/go_waf_sample/application"
	"github.com/ayumu83s/go_waf_sample/controller"
	"github.com/ayumu83s/go_waf_sample/registory"
	"github.com/garyburd/redigo/redis"
)

func main() {
	app := application.Bootstrap("dev")
	fmt.Println("port: ", app.Config.Web.Port)

	defer app.DB.Close() // これをCloseするのはどこが正解なのか・・
	defer app.Redis.Close()

	word, _ := redis.String(app.Redis.Do("GET", "hoge"))
	fmt.Println("word: ", word)
	repoArgs := &registory.RepositoryInput{
		DB: app.DB,
	}
	repo := registory.NewUserRepository(repoArgs)
	controller := controller.NewUser(repo)
	controller.Get(1)

	//controller.Create("EEE", 60)
	controller.Delete(8)
}
