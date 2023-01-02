package main

import "fmt"

func main() {
	var myMap = make(map[string]string)

	// create three keys in the map
	myMap["key1"] = "value1"
	myMap["key2"] = "value2"
	myMap["key3"] = "value3"

	// obtain the number of key/values
	fmt.Println("Map keys:", len(myMap))

	// access the value stored for a key
	fmt.Println(myMap["key1"])

	// iterate over the map with range and print each key/value
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	// delete key3
	delete(myMap, "key3")

	fmt.Println(myMap)
}
