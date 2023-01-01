package main

import (
	"fmt"
)

func SayHello(name string) {
	fmt.Printf("Hello %s\n", name)
	name = "Dave Blogs"
}

func main() {
	name := "Joe Blogs"
	SayHello(name)
	fmt.Println(name)
}
