package main

import (
	"fmt"
)

func main() {
	ping := make(chan struct{}, 1)
	pong := make(chan struct{}, 1)

	ping <- struct{}{}

	for {
		select {
		case <-ping:
			fmt.Println("Ping received")
			pong <- struct{}{}
		case <-pong:
			fmt.Println("Pong received")
		default:
			fmt.Println("Continuing")
			return
		}
	}
}