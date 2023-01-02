package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Index 0:", sl[0])
	fmt.Println("Elements Index 1 - 5 (not including 5):", sl[1:5])
	fmt.Println("Elements Index 5 - end:", sl[5:])
	fmt.Println("Elements Index 0 - 3 (not including 3):", sl[0:3])
}
