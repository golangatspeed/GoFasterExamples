package main

import (
	"fmt"
)

var myVar1 int

var (
	myVar2 int = 1
	myVar3 int
)

var myVar4 int = 4
var myVar5 = 5

func main() {
	var myVar6 int
	var myVar7, myVar8 int
	var myVar9, myVar10 = 9, 10

	myVar11 := 11
	myVar12, myVar13 := 12, 13

	fmt.Println("Go is not so opinionated on some things...")
	fmt.Println(myVar1, myVar2, myVar3, myVar4, myVar5)
	fmt.Println(myVar6, myVar7, myVar8, myVar9, myVar10, myVar11, myVar12, myVar13)
}