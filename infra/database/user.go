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

	row := u.DB.QueryRow("SELECT * FROM user WHERE id = ?", id)

	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.Status,
	)
	if err != nil {
		// not found
		return nil
	}
	return &user
}
