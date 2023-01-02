package main

import "fmt"

func main() {
	var myInt = 128640
	convertedInt := fmt.Sprintf("%d", myInt)

	fmt.Printf("The int is now a string: '%s', of type '%T'\n", convertedInt, convertedInt)
}
