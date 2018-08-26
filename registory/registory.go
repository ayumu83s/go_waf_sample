package registory

// 簡易なDIコンテナ
// リポジトリの実装を返す

import (
	"github.com/ayumu83s/di_sample/domain/repository"
	"github.com/ayumu83s/di_sample/infra/database"
	"github.com/ayumu83s/di_sample/infra/memcache"
)

type Repository interface {
	NewUser() repository.User
	NewUserByCache() repository.User
}

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) NewUser() repository.User {
	return &database.User{}
}

func (r *repositoryImpl) NewUserByCache() repository.User {
	return &memcache.User{}
}
