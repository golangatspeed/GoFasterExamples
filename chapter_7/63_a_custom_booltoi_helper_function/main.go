package main

import "fmt"

func boolToI(b bool) int {
	var bToI int
	if b {
		bToI = 1
	}
	return bToI
}

func main() {
	myBool := true
	myBool2 := false

	boolAsInt := boolToI(myBool)
	boolAsInt2 := boolToI(myBool2)

	fmt.Printf("boolAsInt: %d, boolAsInt2: %d\n", boolAsInt, boolAsInt2)
}