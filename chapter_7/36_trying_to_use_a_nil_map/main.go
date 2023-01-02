package main

import "fmt"

func main() {
	var myMap map[string]string
	//myMap = make(map[string]string)

	myMap["key"] = "value"
	fmt.Println(myMap)
}