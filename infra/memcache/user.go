package memcache

import (
	"fmt"

	"github.com/ayumu83s/go_waf_sample/domain/entity"
)

type User struct {
}

func (u *User) Get(id int) *entity.User {
	fmt.Println("memcache/Get()")
	return &entity.User{
		ID:     id,
		Name:   "aaa",
		Age:    19,
		Status: true,
	}
}

// 定義しないとコンパイルエラー
func (u *User) Create(*entity.User) (*entity.User, error) { return nil, fmt.Errorf("Not Impl") }
func (u *User) Delete(id int) (int64, error)              { return 0, fmt.Errorf("Not Impl") }
