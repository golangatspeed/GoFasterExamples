package main

import (
	"fmt"
)

func main() {
	data := make(chan int, 10)

	fmt.Println("Sending to channel")
	for i := 0; i < 10; i++ {
		data <- i
		if i == 9 {
			fmt.Println("Closing channel")
			close(data)
		}
	}

	fmt.Println("Reading from closed channel")
	for i := 0; i < 10; i++ {
		fmt.Println("Received:", <-data)
	}

	//data <- 10
}