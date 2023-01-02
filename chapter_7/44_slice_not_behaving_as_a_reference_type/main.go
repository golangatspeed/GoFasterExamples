package main

import "fmt"

func addToSlice(sl []int) {
	sl = append(sl, 2)
	fmt.Println(sl)
}

func main() {
	sl := []int{}

	sl = append(sl, 1)
	fmt.Println(sl)

	addToSlice(sl)
	fmt.Println(sl)
}