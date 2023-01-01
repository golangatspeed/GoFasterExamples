

package main

import "fmt"

var myVar = 10

func main() {
	myVar += 5
	var myVar int
	myVar -= 5
	fmt.Println("myVar is:", myVar)
}