package main

import "fmt"

// Named custom struct
type User struct {
	Name string
	Age  int
}

func main() {
	var user User // all fields initialised to nil value
	user.Name = "Joe Blogs"
	user.Age = 20
	fmt.Printf("User is %s. They are %d years of age", user.Name, user.Age)
}
