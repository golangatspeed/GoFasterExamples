package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Printf("G%d is blocked\n", i)
			<-message
			fmt.Printf("G%d is returning\n", i)
		}(i)
	}

	<-time.After(5 * time.Second)

	close(message)

	<-time.After(5 * time.Second)
}