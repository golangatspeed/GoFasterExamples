package main

import "fmt"

func main() {
	activeUser := map[string]bool{
		"Joe Blogs":  true,
		"Dave Blogs": false,
	}

	if active, ok := activeUser["Trevor Blogs"]; !ok {
		fmt.Println("User does not exist")
	} else {
		fmt.Println("Trevor is active:", active)
	}
}
