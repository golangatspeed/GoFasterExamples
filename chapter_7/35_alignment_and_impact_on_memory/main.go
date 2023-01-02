package main

import (
	"fmt"
	"unsafe"
)

type Suboptimal struct {
	bool1 bool
	int1  int
	bool2 bool
	int2  int
	bool3 bool
	int3  int
	bool4 bool
	int4  int
}

type Optimal struct {
	bool1 bool
	bool2 bool
	bool3 bool
	bool4 bool
	int1  int
	int2  int
	int3  int
	int4  int
}

func main() {
	subOptimal := [1000]Suboptimal{}
	sizeSuboptimal := float64(unsafe.Sizeof(subOptimal))

	optimal := [1000]Optimal{}
	sizeOptimal := float64(unsafe.Sizeof(optimal))

	fmt.Printf("Size of suboptimal array = %v bytes\n", sizeSuboptimal)
	fmt.Printf("Size of optimal array = %v bytes\n", sizeOptimal)

	diff := (sizeSuboptimal - sizeOptimal) / sizeSuboptimal * 100
	fmt.Printf("Padding means we're using %d%% more memory...", int(diff))
}