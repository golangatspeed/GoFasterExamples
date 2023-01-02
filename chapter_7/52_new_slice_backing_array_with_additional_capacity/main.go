package main

import (
	"fmt"
)

func main() {
	// amend this start capacity to see the impact
	cp := 100
	sl := make([]int, 0, cp)
	els := 1000
	addr := fmt.Sprintf("%p", sl)

	for i := 0; i < els; i++ {
		sl = append(sl, i)
		if fmt.Sprintf("%p", sl) != addr {

			// when we detect address change inspect
			extra := (float64(cap(sl)) - float64(cp)) / float64(cp) * 100
			fmt.Printf("length: %d, capacity: %d, address: %p. Additional capacity: %0.2f%%\n", len(sl), cap(sl), sl, extra)
			cp = cap(sl)
			addr = fmt.Sprintf("%p", sl)
		}
	}

	overCap := (float64(cap(sl)) - float64(els)) / float64(els) * 100
	fmt.Printf("Backing array over capacity: %0.2f%%", overCap)
}
