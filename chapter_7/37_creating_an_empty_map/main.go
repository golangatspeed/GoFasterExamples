package main

import "fmt"

func main() {
	// declaration followed my initialisation with make
	var myMap map[string]string
	myMap = make(map[string]string)

	// declaration & initialisation with make
	myMap2 := make(map[string]string)

	// declaration & initialisation as empty map literal
	myMap3 := map[string]string{}

	// declaration & initialisation as map literal with key/values
	myMap4 := map[string]string{
		"key": "value",
	}

	fmt.Println(myMap, myMap2, myMap3, myMap4)
}