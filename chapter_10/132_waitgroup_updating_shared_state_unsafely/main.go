package main

import (
	"fmt"
	"sync"
)

func main() {
	var myCounter int

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			myCounter += 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter total is: %d\n", myCounter)
}