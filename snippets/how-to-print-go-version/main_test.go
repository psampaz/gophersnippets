// Runtime: How to print Go version
package main

import (
	"fmt"
	"runtime"
)

func Example() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	// Output:
	// Go version: go1.14.2
}
