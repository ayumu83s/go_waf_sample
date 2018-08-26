package registory

import (
	"github.com/ayumu83s/di_sample/domain/repository"
	"github.com/ayumu83s/di_sample/infra"
)

type Repository interface {
	NewUser() repository.User
}

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) NewUser() repository.User {
	return &infra.User{}
}
