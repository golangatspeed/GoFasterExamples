package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := "hello"
	t := reflect.TypeOf(x)
	fmt.Println(t)
}