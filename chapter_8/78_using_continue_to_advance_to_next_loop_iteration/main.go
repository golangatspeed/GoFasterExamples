package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}

	for _, v := range sl {
		if v%2 == 0 {
			continue
		}
		fmt.Printf("%d is an odd number\n", v)
	}
}