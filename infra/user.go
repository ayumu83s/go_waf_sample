package infra

import "fmt"

type User struct {
}

func (u *User) Get(id int) {
	fmt.Println("infra/Get()")
}
