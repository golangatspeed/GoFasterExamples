package main

import (
	"fmt"
	"reflect"
)

type Numeric interface {
	Signed | Unsigned
}

type Unsigned interface {
	uint8 | uint16 | uint32 | uint64 | uint
}

type Signed interface {
	int8 | int16 | int32 | int64 | int
}

func SumAny[T Numeric](x, y T) T {
	return x + y
}

func main() {
	res := SumAny(int16(1), int16(2))
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))
}
