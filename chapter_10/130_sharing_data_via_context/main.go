package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey string

func main() {
	var name ctxKey = "name"
	ctx := context.WithValue(context.Background(), name, "Joe Blogs")

	// with context.Value()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("My name is %s\n", ctx.Value(ctxKey("name")))
		}
	}()

	time.Sleep(12 * time.Second)
	fmt.Println("Exiting")
}