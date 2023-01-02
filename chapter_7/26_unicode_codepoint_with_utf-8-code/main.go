package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	myString := "🚀"
	fmt.Println("Unicode codepoint represented by rune:", []rune(myString))
	fmt.Println("UTF-8 code represented by up to 4 bytes:", []byte(myString))
	fmt.Println("UTF-8 code represented as Hexadecimal:", hex.EncodeToString([]byte(myString)))
	fmt.Println("length", len(myString))
}
