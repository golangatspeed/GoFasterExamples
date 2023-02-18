package main

import "fmt"

func List(ceil int) []int {
	var res []int
	for i := 0; i < ceil; i++ {
		res = append(res, i)
	}
	return res
}

func PrintList(list []int) {
	fmt.Println(list)
}

func ExampleList() {
	for _, v := range List(5) {
		fmt.Println(v)
	}

	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}

func ExamplePrintList() {
	PrintList([]int{1, 2, 3, 4})

	// Output: 1 2 3 4
}
