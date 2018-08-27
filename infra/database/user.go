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

func (u *User) Create(user *entity.User) (*entity.User, error) {
	res, err := u.DB.Exec("INSERT INTO user (name, age, status) VALUES (?, ?, ?) ",
		user.Name, user.Age, user.Status,
	)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	user.ID = int(id)
	fmt.Println("lastInsertId: ", user.ID)
	return user, nil
}
