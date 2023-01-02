package main

import "fmt"

type Eyes struct {
	Color string
	Shape string
}

type Human struct {
	Eyes
}

type Dog struct {
	Eyes Eyes
}

func main() {
	var human Human
	var dog Dog

	// promoted field access
	human.Color = "Blue"
	human.Shape = "Round"

	// long form field access
	human.Eyes.Color = "Brown"

	// dog.Color = "Brown" // won't build

	fmt.Println(human, dog)
}