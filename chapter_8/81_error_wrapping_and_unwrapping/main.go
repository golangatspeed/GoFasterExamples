package main

import (
	"errors"
	"fmt"
)

func call3() error {
	return fmt.Errorf("call3: this is the original error")
}

func call2() error {
	if err := call3(); err != nil {
		return fmt.Errorf("call2: %w", err)
	}
	return nil
}

func call1() error {
	if err := call2(); err != nil {
		return fmt.Errorf("call1: %w", err)
	}
	return nil
}

func main() {
	if err := call1(); err != nil {
		fmt.Println(err) // prints all wrapped errors

		// unwrap the errors and print
		err1 := errors.Unwrap(err)
		err2 := errors.Unwrap(err1)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}