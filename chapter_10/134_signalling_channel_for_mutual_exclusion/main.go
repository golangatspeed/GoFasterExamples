package main

import (
	"fmt"
	"sync"
)

func main() {
	var myCounter int
	var wg sync.WaitGroup

	var lock = make(chan struct{}, 1)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			lock <- struct{}{}
			myCounter += 1
			<-lock

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter total is: %d\n", myCounter)
}