package main

import "fmt"

func addToSlice(sl []int) []int {
	sl = append(sl, 2)
	return sl
}

func main() {
	sl := []int{}

	sl = append(sl, 1)
	fmt.Println(sl)

	sl = addToSlice(sl)
	fmt.Println(sl)
}