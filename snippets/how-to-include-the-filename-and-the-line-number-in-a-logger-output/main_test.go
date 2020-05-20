// Logging: How to include the filename and the line number in a logger output
package main

import (
	"log"
	"os"
)

func Example() {
	// You can configure the output format of a logger using
	// flags https://golang.org/pkg/log/#pkg-constants
	// Use log.LshortfileIf to print the filename without
	// the full path and the line number or
	// Use log.Llongfile to print the filename with
	// the full path and the line number
	// The log.Llongfile example is not included below
	// because playground use different path each time
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Hello world")
	// Output:
	// prog.go:20: Hello world
}
