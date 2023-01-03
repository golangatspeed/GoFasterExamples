package main

import (
	"fmt"
)

func main() {
	myName := "Joe Blogs"

	switch myName {
	case "Joe Blogs":
		fmt.Println("Hi Joe!")
	case "Dave Blogs":
		fmt.Println("Hi Dave!")
	default:
		fmt.Println("Hi there!")
	}
}
