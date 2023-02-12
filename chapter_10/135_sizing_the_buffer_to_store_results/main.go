package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	store := make(chan int, 3)

	// fill the channel buffer
	fmt.Println("Filling the buffer")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			store <- i
			fmt.Printf("Channel buffer contains %d values\n", len(store))
			wg.Done()
		}(i)
	}

	wg.Wait()

	// read back from the channel buffer
	fmt.Println("Reading the buffer")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("Reading value %d from channel buffer\n", <-store)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Finished")
}