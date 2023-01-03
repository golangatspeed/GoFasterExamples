package main

import (
	"fmt"
)

func anonymousReturns(firstname, lastname string) string {
	return fmt.Sprintf("%s %s", firstname, lastname)
}

func namedReturns(firstname, lastname string) (fullname string) {
	fullname = fmt.Sprintf("%s %s", firstname, lastname)
	return
}

func main() {
	fmt.Println(anonymousReturns("Joe", "Bloggs"))
	fmt.Println(namedReturns("Joe", "Bloggs"))
}