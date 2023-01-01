package main

import (
	"fmt"
)

const (
	const1 = iota
	_
	_
	const4
	const5
)

func main() {
	fmt.Println("constants values are:", const1, const4, const5)
}
