package main

import (
	"fmt"
)

func Sender(ch chan<- string, message string) {
	// chan<- unidirectional send
	ch <- message
}

func Receiver(ch <-chan string) {
	// <-chan unidirectional receive
	fmt.Println(<-ch)
}

func SenderReceiver(ch chan string, message string) {
	// chan bi-directional send or receive
	ch <- message
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string, 1)

	Sender(ch, "Hello World")
	Receiver(ch)
	SenderReceiver(ch, "Hello World again")
}