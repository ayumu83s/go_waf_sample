package memcache

import (
	"fmt"
)

type User struct {
}

func (u *User) Get(id int) {
	fmt.Println("memcache/Get()")
}
