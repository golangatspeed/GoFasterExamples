package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	f, _ := os.Create("cpu.out")
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// rest of the program
}