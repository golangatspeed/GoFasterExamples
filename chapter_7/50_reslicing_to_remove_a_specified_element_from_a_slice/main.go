package main

import "fmt"

func removeAtIndex(sl []int, index int) []int {
	return append(sl[:index], sl[index+1:]...)
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sl)
	sl = removeAtIndex(sl, 9)
	fmt.Println(sl)
}
