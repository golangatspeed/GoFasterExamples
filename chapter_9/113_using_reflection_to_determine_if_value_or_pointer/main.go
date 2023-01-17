package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := "hello" // value
	y := &x // pointer

	v1 := reflect.ValueOf(x)
	fmt.Println("x is pointer:", v1.Kind() == reflect.Ptr)

	v2 := reflect.ValueOf(y)
	fmt.Println("y is a pointer:", v2.Kind() == reflect.Ptr)
}