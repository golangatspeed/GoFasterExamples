package main

import (
	"errors"
	"fmt"
)

func dummyFunctionError() {
	err := errors.New("dummy unhandlable error")
	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	dummyFunctionError()
}
