// go run -gcflags="-m" ./main.go
// go test -bench . -benchmem
package main

type Customer struct {
	Name  string
	Email string
	Age   int
}

//go:noinline
func NewCustomer(name string) Customer {
	cust := Customer{Name: name}
	return cust
}

//go:noinline
func NewCustomer2(name string) *Customer {
	cust2 := Customer{Name: name}
	return &cust2
}

//go:noinline
func NewCustomer3(name string) *Customer {
	return &Customer{Name: name}
}

func main() {
	input := "Joe Blogs"
	_ = NewCustomer(input)
	//_ = NewCustomer2(input)
	//_ = NewCustomer3(input)
}