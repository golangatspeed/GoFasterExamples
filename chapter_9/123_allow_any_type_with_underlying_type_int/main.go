package main

import (
	"fmt"
	"reflect"
)

type OnlyInt interface {
	int
	// ~int
}

type myInt int

func SumAny[T OnlyInt](x, y T) T {
	return x + y
}

func main() {
	var myInt myInt = 1
	res := SumAny(myInt, 2)
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))
}