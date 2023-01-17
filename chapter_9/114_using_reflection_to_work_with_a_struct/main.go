package main

import (
	"fmt"
	"reflect"
)

type Person struct{
	Age int
}

func (m *Person) SayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func main() {
	p := &Person{Age:22}
	v := reflect.ValueOf(p)

	// call the receiver/method with argument
	sayHello := v.MethodByName("SayHello")
	name := reflect.ValueOf("Joe Blogs")
	sayHello.Call([]reflect.Value{reflect.Value(name)})

	// access the Age property, and change it
	age := v.Elem().FieldByName("Age")
	fmt.Println("Current age:", age.Int())
	age.SetInt(42)
	fmt.Println("New age:", age.Int())
}