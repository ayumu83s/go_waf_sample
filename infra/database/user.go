package database

import (
	"database/sql"
	"fmt"

	"github.com/ayumu83s/go_waf_sample/domain/entity"
)

type User struct {
	DB *sql.DB
}

func (u *User) Get(id int) *entity.User {
	fmt.Println("database/Get()")
	return &entity.User{
		ID:     id,
		Name:   "aaa",
		Age:    30,
		Status: false,
	}
}
