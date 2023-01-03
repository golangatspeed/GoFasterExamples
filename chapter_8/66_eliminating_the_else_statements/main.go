package main

import "fmt"

func isEvenLessThan(n int, c int) (bool, bool) {
	var even, less bool // initialised to false

	if n%2 == 0 {
		even = true
	}
	if n < c {
		less = true
	}

	return even, less
}

func main() {
	fmt.Println(isEvenLessThan(2, 10))
}