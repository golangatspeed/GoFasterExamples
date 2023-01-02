package main

import "fmt"

func main() {
	var sl1 []int
	sl2 := []int{}
	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5)
	sl5 := make([]int, 1, 4)

	// check length
	fmt.Println(len(sl1), len(sl2), len(sl3), len(sl4), len(sl5))

	// check capacity
	fmt.Println(cap(sl1), cap(sl2), cap(sl3), cap(sl4), cap(sl5))

	if sl1 == nil {
		fmt.Println("sl1 is nil!")
	}
	if sl2 != nil && sl3 != nil && sl4 != nil && sl5 != nil {
		fmt.Println("sl2, sl3, sl4 and sl5 are not nil!")
	}
}
