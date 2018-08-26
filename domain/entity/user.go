package entity

type User struct {
	ID     int
	Name   string
	Age    int
	Status bool
}

func (u *User) IsAdult() bool {
	if u.Age >= 20 {
		return true
	}
	return false
}
