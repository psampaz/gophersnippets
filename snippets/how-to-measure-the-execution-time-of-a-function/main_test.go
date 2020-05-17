// Time: How to measure the execution time of a function
//
// This example demonstrates simple execution time calculation
// If you need a detailed report on a function's performance
// you should use Go benchmark functions
// https://golang.org/pkg/testing/#hdr-Benchmarks
// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
package main

import (
	"fmt"

	"time"
)

func f() {
	// first we calculate the exact time the function
	// started the execution
	start := time.Now()
	// Then we calculate the execution time duration
	// inside a deferred function, since it will be execute after f() returns
	defer func(start time.Time) {
		dur := time.Since(start)
		fmt.Printf("f() took %0.0f to execute", dur.Seconds())
	}(start)
	// pause the execution for 2 seconds
	time.Sleep(2 * time.Second)
}

func Example() {
	f()
	// Output:
	// f() took 2 to execute
}
