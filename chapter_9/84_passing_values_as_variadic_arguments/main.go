package main

import "fmt"

func myVariadicFunc(name string, address ...string) {
	fmt.Printf("Hello %s!\n", name)
	fmt.Println("Addresses:")
	if len(address) > 0 {
		for i, addr := range address {
			fmt.Printf("%d: %s\n", i+1, addr)
		}
	} else {
		fmt.Println("No address data supplied")
	}
}

func main() {
	// single argument
	fmt.Println("Single Argument")
	myVariadicFunc("Joe Bloggs", "Address 1")

	// multiple arguments
	fmt.Println("\nMultiple Arguments")
	myVariadicFunc("Joe Bloggs", "Address 1", "Address 2", "Address 3")

	// no argument
	fmt.Println("\nNo argument")
	myVariadicFunc("Joe Bloggs")

	// passing a pre-built slice, note the trailing spread operator
	fmt.Println("\nPassing a slice as a variadic argument")
	addresses := []string{"Address 1", "Address 2"}
	myVariadicFunc("Joe Bloggs", addresses...)
}