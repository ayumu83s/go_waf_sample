package repository

import "github.com/ayumu83s/go_waf_sample/domain/entity"

// CRUDに関するI/Fを定義

type User interface {
	Get(id int) *entity.User
	Create(*entity.User) (*entity.User, error)
	Delete(id int) (int64, error)
}
