package main

import "fmt"

type numer int
type denom int

func divider(d denom, n numer) float64 {
	return float64(n) / float64(d)
}

func main() {
	var n numer = 1
	var d denom = 4

	result := divider(n, d)
	fmt.Println(result)
}
