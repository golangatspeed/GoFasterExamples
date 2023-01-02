package main

import (
	"fmt"
	"strconv"
)

func main() {
	myInt := 10
	myIntString := "10"
	myFloatString := "3.1415926535"

	// int to string
	resStr := strconv.Itoa(myInt)
	fmt.Printf("resStr is '%T, of '%v'\n", resStr, resStr)

	// string to int
	if resInt, err := strconv.Atoi(myIntString); err == nil {
		fmt.Printf("resInt is '%T', '%v'\n", resInt, resInt)
	}

	// string to float
	if resFloat, err := strconv.ParseFloat(myFloatString, 64); err == nil {
		fmt.Printf("resFloat is '%T', '%v'\n", resFloat, resFloat)
	}
}
