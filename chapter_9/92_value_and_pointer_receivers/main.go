package main

import "fmt"

type customer struct {
	Name string
}

// UpdateName is a pointer receiver
func (c *customer) UpdateName(newStr string) {
	c.Name = newStr
}

// PrintName is a value receiver
func (c customer) PrintName() {
	fmt.Printf("This is reading the value: %s\n", c.Name)
	c.PrintLine() // calling another receiver
}

func (customer) PrintLine() {
	fmt.Println("-------------------------------------")
}

func main() {
	var cust = &customer{"Joe Blogs"}

	// read via value receiver
	cust.PrintName()

	// update
	cust.UpdateName("Dave Blogs")
	fmt.Printf("Updated string is value: %s\n", cust.Name)
	cust.PrintLine()
}