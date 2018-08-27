package service

import (
	"fmt"

	"github.com/ayumu83s/go_waf_sample/domain/repository"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(u repository.User) *UserService {
	return &UserService{u}
}

func (u *UserService) Get(id int) {
	user := u.userRepo.Get(id)
	if user != nil {
		fmt.Println(user.IsAdult())
	}
}
