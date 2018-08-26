package controller

import (
	"github.com/ayumu83s/di_sample/domain/service"
	"github.com/ayumu83s/di_sample/registory"
)

type User struct {
	repo registory.UserRepository
}

func NewUser(repo registory.UserRepository) *User {
	return &User{repo}
}

func (u *User) Get(id int) {
	repo := u.repo.NewUser()
	//repo := u.repo.NewUserByCache()
	userService := service.NewUserService(repo)
	userService.Get(id)
}
