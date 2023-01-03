package main

import (
	"fmt"
)

func main() {
	switch letter := "z"; letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("letter was a vowel")
	default:
		fmt.Println("letter was not a vowel")
	}
}
