package main

import "fmt"

// Named custom struct
type User struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	user := User{
		"Joe Blogs",
		20,
		"Male",
	}
	fmt.Println(user)
}
