package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	var results [30]int

	for i := 1; i <= 30; i++ {
		results[i-1] = i * 10
	}

	fmt.Println(results)
	fmt.Println("Completed in", time.Since(t))
}