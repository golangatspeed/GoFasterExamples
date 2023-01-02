package main

import "fmt"

func send(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func read(ch chan int) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func main() {
	ch := make(chan int)
	fmt.Println(ch)
	go read(ch)
	send(ch)
}
