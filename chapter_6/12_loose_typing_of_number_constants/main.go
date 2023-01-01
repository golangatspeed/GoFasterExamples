package main

import (
	"fmt"
)

const myConst = 1
var myVar = 1

func needsAFloat(val float64) {
	fmt.Println("looks like we got a float:", val)
}

func main() {
	fmt.Printf("myConst is type '%T'\n", myConst)
	fmt.Printf("myVar is type '%T'\n", myVar)
	needsAFloat(myConst)
	//needsAFloat(myVar)
}
