package main

import (
	"fmt"
	_ "github.com/golangatspeed/pkg/sideeffect"
)

func init() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("Hello World again!")
}