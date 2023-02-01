package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// wait 5 seconds then call the cancelFunc
	// which will stop any running goroutines which honour context
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	// no context
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("I'm Dave Blogs and I AM unstoppable")
		}
	}()

	// with context
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("I'm Joe Blogs I'm unstoppable!")
			case <-ctx.Done():
				err := ctx.Err()
				fmt.Printf("I SPOKE TOO SOON... '%v' OH NO!\n", err)
				fmt.Println("Cleanup and graceful shutdown type stuff performed here...")
				return
			}
		}
	}()

	// exit after 20 seconds
	time.Sleep(20 * time.Second)
}