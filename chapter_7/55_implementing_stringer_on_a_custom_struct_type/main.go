package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("Hey there I'm %v and I'm %v years young!", p.name, p.age)
}

func main() {
	st := person{
		"Joe Blogs",
		25}
	fmt.Println(st)
}
