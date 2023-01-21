package main

import "fmt"

func sumInt64(x, y int64) int64 {
	return x + y
}

func main() {
	fmt.Println(sumInt64(1, 2))
}