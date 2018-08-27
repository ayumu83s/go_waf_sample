package service

import (
	"fmt"

	"github.com/ayumu83s/go_waf_sample/domain/entity"
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

func (u *UserService) Create(name string, age int) {
	input := &entity.User{
		Name:   name,
		Age:    age,
		Status: true,
	}
	user, err := u.userRepo.Create(input)
	if err != nil {
		return
	}
	fmt.Println(user.ID)
	fmt.Println(user.IsAdult())
}

func (u *UserService) Delete(id int) {
	u.userRepo.Delete(id)
}
