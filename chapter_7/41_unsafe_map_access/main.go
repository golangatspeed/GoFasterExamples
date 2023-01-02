package main

import "fmt"

func main() {
	activeUser := map[string]bool{
		"Joe Blogs":  true,
		"Dave Blogs": false,
	}

	// keys exists no issue
	fmt.Println("Joe is active:", activeUser["Joe Blogs"])
	fmt.Println("Dave is active:", activeUser["Dave Blogs"])
	// key does not exist
	fmt.Println("Trevor is active:", activeUser["Trevor Blogs"])
}
