package main

import (
	"fmt"
	"time"
)

func main() {

START:

	fmt.Println("Timer running")
	time.Sleep(3 * time.Second)
	goto START
}