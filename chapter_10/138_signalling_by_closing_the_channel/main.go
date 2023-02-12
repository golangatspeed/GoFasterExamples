package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		for i := 5; i > 0; i-- {
			fmt.Printf("Signalling done in %d seconds\n", i)
			time.Sleep(1 * time.Second)
		}

		close(done)
	}()

	<-done

	fmt.Println("OK, exiting")
}
