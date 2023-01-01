package stringutil_test

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func ExampleReverse() {
	fmt.Println(stringutil.Reverse("hello"))
	// Output: olleh
}