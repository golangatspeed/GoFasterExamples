package main

func main() {
	var i interface{}
	i = "this is a string"

	// we assert that it is in interface contains a integer value
	x := i.(int)
	_ = x
}