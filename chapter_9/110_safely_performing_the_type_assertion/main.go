package main

import "fmt"

func main() {
	var i interface{}
	i = "this is a string"

	// we assert that it is in interface contains an integer value
	x, ok := i.(int)
	if ok {
		fmt.Printf("We have an integer of: %d", x)
	} else {
		fmt.Printf("We don't have an integer, setting to zero value: %d", x)
	}
}