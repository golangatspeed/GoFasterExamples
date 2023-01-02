package main

import "fmt"

func main() {
	intArray := [3]int{1, 2, 3}
	intArray2 := [...]int{4, 5, 6}

	fmt.Println(intArray)
	fmt.Println(intArray2)
}