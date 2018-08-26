package repository

import "github.com/ayumu83s/di_sample/domain/entity"

// CRUDに関するI/Fを定義

type User interface {
	Get(id int) *entity.User
}
