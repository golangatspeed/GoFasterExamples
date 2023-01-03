package main

import (
	"fmt"
)

func anonymousReturns() (string, string) {
	var firstname, lastname string // we had to declare lastname
	firstname = "Joe"
	return firstname, lastname // we had to return lastname
}

func namedReturns() (firstname, lastname string) {
	firstname = "Joe"
	return // did we forget the signature also includes `lastname`?
}

func main() {
	mask := "firstname: '%s', lastname: '%s'\n"
	f1, l1 := anonymousReturns()
	fmt.Printf(mask, f1, l1)
	f2, l2 := namedReturns()
	fmt.Printf(mask, f2, l2)
}