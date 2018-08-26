package registory

// 簡易なDIコンテナ
// リポジトリの実装を返す

import (
	"github.com/ayumu83s/go_waf_sample/domain/repository"
	"github.com/ayumu83s/go_waf_sample/infra/database"
	"github.com/ayumu83s/go_waf_sample/infra/memcache"
)

type UserRepository interface {
	NewUser() repository.User
	NewUserByCache() repository.User
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) NewUser() repository.User {
	return &database.User{}
}

func (r *userRepositoryImpl) NewUserByCache() repository.User {
	return &memcache.User{}
}
