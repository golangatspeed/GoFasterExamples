package main

import "fmt"

// Named custom struct
type User struct {
	Name string
	Age  int
}

// Second named custom struct
type User2 struct {
	Name string
	Age  int
}

func main() {
	var user User

	// Anonymous struct created at point of use
	data := struct {
		Name string
		Age  int
	}{
		Name: "Joe Blogs",
		Age:  20,
	}

	user = data

	//var user2 User2
	//user2 = user

	fmt.Println(user)
}