package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl2 := sl[1:5]

	fmt.Printf("len sl: %d , cap sl: %d\n", len(sl), cap(sl))
	fmt.Printf("len sl2: %d , cap sl2: %d\n", len(sl2), cap(sl2))

	sl2[2] = 20
	fmt.Println("Slice sl2:", sl2)
	fmt.Println("Slice sl:", sl)
}
