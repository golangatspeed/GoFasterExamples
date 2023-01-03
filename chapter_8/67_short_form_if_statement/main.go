package main

import "fmt"

func isEvenLessThan(n int, c int) (bool, bool) {
	var even, less bool

	if n%2 == 0 {
		even = true
	}
	if n < c {
		less = true
	}

	return even, less
}

func main() {
	if even, _ := isEvenLessThan(2, 10); even {
		fmt.Println("Result was even")
	}

	//fmt.Println(even)
}
