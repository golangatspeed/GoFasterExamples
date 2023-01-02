package main

import "fmt"

func alterMap(mp *map[string]string) {
	(*mp)["myKey"] = "my new value"
}

func main() {
	myMap := map[string]string{
		"myKey": "my value",
	}

	fmt.Println(myMap)
	alterMap(&myMap)
	fmt.Println(myMap)
}
