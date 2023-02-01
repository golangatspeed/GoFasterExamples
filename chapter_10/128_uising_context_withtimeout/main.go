package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// with context
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("I'm Joe Blogs I've got 10 seconds to talk")
			case <-ctx.Done():
				err := ctx.Err()
				fmt.Printf("Got to go now... '%v'\n", err)
				fmt.Println("Cleanup e.t.c")
				return
			}
		}
	}()

	time.Sleep(12 * time.Second)
	fmt.Println("Exiting")
}