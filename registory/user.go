package registory

// 簡易なDIコンテナ
// リポジトリの実装を返す

import (
	"database/sql"

	"github.com/ayumu83s/go_waf_sample/domain/repository"
	"github.com/ayumu83s/go_waf_sample/infra/database"
	"github.com/ayumu83s/go_waf_sample/infra/memcache"
)

type UserRepository interface {
	NewUser() repository.User
	NewUserByCache() repository.User
}

type userRepositoryImpl struct {
	DB *sql.DB
}

type RepositoryInput struct {
	DB *sql.DB
}

func NewUserRepository(input *RepositoryInput) UserRepository {
	return &userRepositoryImpl{
		DB: input.DB,
	}
}

func (r *userRepositoryImpl) NewUser() repository.User {
	return &database.User{
		DB: r.DB,
	}
}

func (r *userRepositoryImpl) NewUserByCache() repository.User {
	return &memcache.User{}
}
