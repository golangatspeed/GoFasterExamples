package main

import "fmt"

func incrementElement(sl []int) {
	for i := range sl {
		sl[i]++
	}
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl)

	incrementElement(sl)
	fmt.Println(sl)
}
