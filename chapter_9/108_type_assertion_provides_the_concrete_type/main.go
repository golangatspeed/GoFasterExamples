package main

import "fmt"

type MyInterface interface {
	Do() error
}

type MyStruct struct {
	Greeting string
}

func (ms MyStruct) Do() error {
	return nil
}

func main() {
	var x MyInterface
	x = MyStruct{Greeting: "Hello"}

	// cannot access the concrete type, `Greeting` property is not visible
	fmt.Printf("This will generate  compilation error", x.Greeting)

	// perform type assertion
	y := x.(MyStruct)

	// we can now access the structs properties
	fmt.Printf("'Greeting' set to: %s\n", y.Greeting)
}