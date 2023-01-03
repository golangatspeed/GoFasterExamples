package main

import (
	"fmt"
	"strings"
)

// Formatter accepts any func of type func (string) string
func Formatter(str string, format func(in string) string) string {
	return format(str)
}

func main() {
	upper := func(in string) string {
		return strings.ToUpper(in)
	}

	lower := func(in string) string {
		return strings.ToLower(in)
	}

	myText := "SomE rANDOm text TO foRmat"
	fmt.Println(Formatter(myText, upper))
	fmt.Println(Formatter(myText, lower))
}