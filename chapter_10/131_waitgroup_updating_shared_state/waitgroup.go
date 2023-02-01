package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	var results [30]int

	var wg sync.WaitGroup
	//wg.Add(29)

	for i := 1; i <= 30; i++ {
		wg.Add(1)
		go func(i int) {
			results[i-1] = i * 10
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(results)
	fmt.Println("Completed in", time.Since(t))
}