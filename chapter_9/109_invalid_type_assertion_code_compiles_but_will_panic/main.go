package main

func main() {
	var i interface{}
	i = "this is a string"

	// we assert the interface contains an integer value
	x := i.(int)
	_ = x
}