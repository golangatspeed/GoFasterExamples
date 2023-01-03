package main

import (
	"fmt"
	"time"
)

type MyError struct {
	// struct properties here, if any
}

func (e MyError) Error() string {
	return fmt.Sprintf("this is custom error generated at %s", time.Now())
}

func dummyFunctionError() error {
	return &MyError{}
}

func main() {
	if err := dummyFunctionError(); err != nil {
		fmt.Println(err)
	}
}