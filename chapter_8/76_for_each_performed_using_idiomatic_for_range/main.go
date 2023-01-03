package main

import "fmt"

func main() {
	sl := []int{1, 2}
	mp := map[string]string{"key1": "value1", "key2": "value2"}

	// slice/array
	for index, value := range sl {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// map
	for key, value := range mp {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}