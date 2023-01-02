package main

import "fmt"

func divider(d, n int) float64 {
	return float64(n) / float64(d)
}

func main() {
	var n = 1
	var d = 4

	result := divider(n, d)
	fmt.Println(result) // we think we will get 0.25
}
