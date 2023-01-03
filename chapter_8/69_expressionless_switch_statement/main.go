package main

import (
	"fmt"
)

func main() {
	num := 12

	switch {
	case num >= 0 && num <= 10:
		fmt.Println("Between 0 and 10")
	case num >= 10:
		fmt.Println("Greater than 10")
	}
}