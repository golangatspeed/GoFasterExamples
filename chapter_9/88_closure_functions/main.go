package main

import "fmt"

func main() {
	x := 10

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment()) // prints 11
	fmt.Println(increment()) // prints 12
}
