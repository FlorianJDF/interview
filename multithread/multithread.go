package main

import "fmt"

type User struct {
	ID int
}

func (s *User) ChangeID(id int) {
	s.ID = id
}

func main() {
	var ids []int

	for i := 1; i < 100; i++ {
		ids = append(ids, i)
	}

	u := User{0}
	for _, id := range ids {
		go u.ChangeID(id)
		fmt.Println(u.ID)
	}
}
