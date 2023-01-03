package main

import "fmt"

func isEvenLessThan(n int, c int) (bool, bool) {
	if n%2 == 0 {
		if n < c {
			return true, true
		} else {
			return true, false
		}
	} else {
		if n < c {
			return false, true
		} else {
			return false, false
		}
	}
}

func main() {
	fmt.Println(isEvenLessThan(2, 10))
}
