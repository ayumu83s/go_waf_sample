package controller

import (
	"github.com/ayumu83s/di_sample/registory"
	"github.com/ayumu83s/di_sample/service"
)

type User struct {
	repo registory.Repository
}

func NewUser(repo registory.Repository) *User {
	return &User{repo}
}

func (u *User) Get(id int) {
	repo := u.repo.NewUser()
	userService := service.NewUser(repo)
	userService.Get(id)
}
