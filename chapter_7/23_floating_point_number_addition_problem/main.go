package main

import (
	"fmt"
)

func main() {

	a := 2.1
	b := 4.2

	c := a + b
	d := 2.1 + 4.2
	e := float64(2.1) + float64(4.2)

	fmt.Println(c == d)
	fmt.Println(c == e)
}