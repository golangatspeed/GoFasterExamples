package main

import (
	"fmt"
	"reflect"
	"time"
)

var i int = 42

func DirectAccess() time.Duration {
	start := time.Now()
	for j := 0; j < 100000; j++ {
		val := i
		_ = val
	}
	elapsed := time.Since(start)
	fmt.Printf("Direct variable access: %s\n", elapsed)
	return elapsed
}

func ReflectAccess() time.Duration {
	start := time.Now()
	for j := 0; j < 100000; j++ {
		val := reflect.ValueOf(i)
		_ = val
	}
	elapsed := time.Since(start)
	fmt.Printf("Access using reflection: %s\n", elapsed)
	return elapsed
}

func main() {
	elapsed1 := DirectAccess()
	elapsed2 := ReflectAccess()
	inc := ((elapsed2 - elapsed1) / elapsed1) * 100
	fmt.Printf("Reflection takes %d%% longer than direct access", inc)
}