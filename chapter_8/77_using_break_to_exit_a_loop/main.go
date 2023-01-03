package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	found := 0

	for _, v := range sl {
		if v == 3 {
			found = v
			break
		}
	}

	fmt.Println("found:", found)
}