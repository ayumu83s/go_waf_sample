package database

import (
	"fmt"

	"github.com/ayumu83s/di_sample/domain/entity"
)

type User struct {
}

func (u *User) Get(id int) *entity.User {
	fmt.Println("database/Get()")
	return &entity.User{
		ID:     id,
		Name:   "aaa",
		Age:    20,
		Status: false,
	}
}
