package main

import (
	"fmt"
)

const (
	blue = iota
	//red
	yellow
	purple
)

func main() {
	chosenColor := yellow
	switch chosenColor {
	case 0:
		fmt.Println("Chose blue")
	case 1:
		fmt.Println("Chose yellow")
	case 2:
		fmt.Println("Chose purple")
	}
}
