package main

import (
	"github.com/ayumu83s/di_sample/controller"
	"github.com/ayumu83s/di_sample/registory"
)

func main() {
	repo := registory.NewRepository()
	controller := controller.NewUser(repo)
	controller.Get(1)
}
