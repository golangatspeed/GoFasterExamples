package main

import "fmt"

type numer int

func (n numer) String() string {
	return fmt.Sprintf("Numerator is set to %d", n)
}

func main() {
	var n numer = 1
	fmt.Println(n)
}
