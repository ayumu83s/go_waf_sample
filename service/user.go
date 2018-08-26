package service

import "github.com/ayumu83s/di_sample/repository"

type User struct {
	userRepo repository.User
}

func NewUser(u repository.User) *User {
	return &User{u}
}

func (u *User) Get(id int) {
	u.userRepo.Get(id)
}
