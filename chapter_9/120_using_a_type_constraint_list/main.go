package main

import (
	"fmt"
	"reflect"
)

func SumAny[T int8 | int32 | int64 | int](x, y T) T {
	return x + y
}

func main() {
	res := SumAny(int16(1), int16(2))
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))
}