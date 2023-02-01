package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Will this print?")
	}()

	//time.Sleep(1 * time.Second)
}