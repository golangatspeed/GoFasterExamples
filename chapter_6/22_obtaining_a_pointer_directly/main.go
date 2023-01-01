package main

import "fmt"

type User struct {
	Name string
}

func ChangeName(name string, user *User) {
	user.Name = name
}

func main() {
	u1 := new(User)
	u2 := &User{
		Name: "Joe Blogs",
	}

	ChangeName("Dave Blogs", u1)
	ChangeName("Trevor Blogs", u2)

	fmt.Println("u1:", u1)
	fmt.Println("u2:", u2)
}
