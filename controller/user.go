package controller

import (
	"github.com/ayumu83s/di_sample/domain/service"
	"github.com/ayumu83s/di_sample/registory"
)

type User struct {
	repo registory.Repository
}

func NewUser(repo registory.Repository) *User {
	return &User{repo}
}

func (u *User) Get(id int) {
	repo := u.repo.NewUser()
	userService := service.NewUserService(repo)
	userService.Get(id)
}
