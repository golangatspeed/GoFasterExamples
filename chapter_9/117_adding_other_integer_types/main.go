package main

import "fmt"

func sumInt8(x, y int8) int8 {
	return x + y
}

func sumInt16(x, y int16) int16 {
	return x + y
}

func sumInt32(x, y int32) int32 {
	return x + y
}

func sumInt64(x, y int64) int64 {
	return x + y
}

func main() {
	fmt.Println(sumInt8(1, 2))
	fmt.Println(sumInt16(1, 2))
	fmt.Println(sumInt32(1, 2))
	fmt.Println(sumInt64(1, 2))
}