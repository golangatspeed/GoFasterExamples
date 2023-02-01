package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Setting GOMAXPROCS to 5, previously was %d\n", runtime.GOMAXPROCS(5))
	fmt.Printf("Setting GOMAXPROCS to 2, previously was %d\n", runtime.GOMAXPROCS(2))
}
