package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	deadline := time.Now().Add(10 * time.Second) // 10 seconds in future
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	// with context
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Printf("I'm Joe Blogs I can talk until %s\n", deadline)
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