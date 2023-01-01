package main

import (
	"fmt"
)

var myVar = 10

func main() {
	// global scope
	fmt.Println(myVar)

	myVar := 5

	for i := 0; i < 1; i++ {
		//control structure scope
		myVar := 0
		fmt.Println(myVar)
	}

	// function scope
	fmt.Println(myVar)
}
