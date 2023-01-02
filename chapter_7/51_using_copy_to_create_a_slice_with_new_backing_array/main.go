package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl2 := make([]int, 4)

	// copy parameter order is the destination then the source
	copy(sl2, sl[1:5])

	sl2[2] = 20
	fmt.Println("Slice sl2:", sl2)
	fmt.Println("Slice sl:", sl)
}
