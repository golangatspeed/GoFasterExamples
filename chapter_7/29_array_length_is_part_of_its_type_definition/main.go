package main

import "fmt"

func main() {
	var intArray [5]int
	var intArray2 [6]int

	if intArray == intArray2 {
		// this will never print
		fmt.Println("Array values and type are same")
	}
}
