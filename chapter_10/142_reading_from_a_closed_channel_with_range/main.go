package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int, 10)

	// worker
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending data")
			data <- i
		}
		close(data)
		fmt.Println("Channel closed")
	}()

	// handler
	for d := range data {
		fmt.Printf("Receiving %d\n", d)
		time.Sleep(1 * time.Second)
	}
}