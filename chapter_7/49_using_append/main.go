package main

import "fmt"

func main() {
	sl := []int{1, 2}
	sl2 := []int{7, 8, 9}

	sl = append(sl, 3)       // add single element
	sl = append(sl, 4, 5, 6) // add many new elements
	sl = append(sl, sl2...)  // append all of sl2 using the spread operator

	fmt.Println(sl)
}
