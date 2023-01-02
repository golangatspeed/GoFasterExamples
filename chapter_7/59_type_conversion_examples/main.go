package main

import "fmt"

func main() {
	var myInt = 1
	var myFloat = 0.0
	var myString = "Hello World"
	var myBytes = []byte{240, 159, 154, 128}

	fmt.Printf("Before conversion, myInt is: %T, myFloat is: %T, myString is: %T, myBytes is: %T\n",
		myInt, myFloat, myString, myBytes)

	convertedInt := float64(myInt)
	convertedFloat := int(myFloat)
	convertedBytes := string(myBytes)
	convertedString := []byte(myString)

	fmt.Printf("After conversion, myInt is: %T, myFloat is: %T, myString is: %T, myBytes is: %T\n",
		convertedInt, convertedFloat, convertedString, convertedBytes)

	// incompatible data types
	//convertedFloat2 := string(myFloat)
	//convertedBool := int(true)
	//convertedString2 := int(myString)
}