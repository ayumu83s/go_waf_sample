package service

import "github.com/ayumu83s/di_sample/domain/repository"

type UserService struct {
	userRepo repository.User
}

func NewUserService(u repository.User) *UserService {
	return &UserService{u}
}

func (u *UserService) Get(id int) {
	u.userRepo.Get(id)
}