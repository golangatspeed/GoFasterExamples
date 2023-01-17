package main

import "fmt"

func main() {
	var i interface{} = "hello"

	switch v := i.(type) {
	case int:
		fmt.Printf("i is an int with value %v\n", v)
	case string:
		fmt.Printf("i is a string with value %v\n", v)
	default:
		fmt.Printf("i is of an unknown type with value %v\n", v)
	}
}