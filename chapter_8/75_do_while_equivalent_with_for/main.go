package main

import "fmt"

func main() {
	n := 1
	for ok := true; ok; ok = n != 5 {
		fmt.Println("Iteration", n)
		n++
	}
	fmt.Printf("n finished at %d\n", n)
}