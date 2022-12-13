package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func (s *User) ChangeName(name string) {
	s.Name = name
}

func main() {
	var names = []string{"a", "b", "c"}

	u := User{1, "foo"}
	for _, name := range names {
		go u.ChangeName(name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
