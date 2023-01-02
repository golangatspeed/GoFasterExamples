package main

import (
	"fmt"
)

func main() {
	var intArray [5]int

	fmt.Println("Array length:", len(intArray))
	fmt.Println("Array capacity:", cap(intArray))

	for i := range intArray {
		fmt.Println("Array element index", i, "contains value", intArray[i])
	}

	intArray[0] = 1
	fmt.Println(intArray)
}