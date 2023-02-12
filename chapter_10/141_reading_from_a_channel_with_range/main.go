package main

import (
	"fmt"
)

func main() {
	data := make(chan int)

	// worker
	go func() {
		for i := 0; i < 10; i++ {
			data <- i
		}
		close(data)
		fmt.Println("Channel closed")
	}()

	// handler
	for d := range data {
		fmt.Printf("Receiving %d\n", d)
	}
}