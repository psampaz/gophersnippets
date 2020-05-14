// Runtime: How to print information about the operating system, architecture and pointer size
package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func Example() {
	fmt.Printf("running program's operating system target: %s\n", runtime.GOOS)
	fmt.Printf("running program's architecture target: %s\n", runtime.GOARCH)
	var ptr uintptr
	ptrSize := int(unsafe.Sizeof(ptr))
	fmt.Printf("pointer size: %d\n", ptrSize)
	// Output:
	// running program's operating system target: linux
	// running program's architecture target: amd64
	// pointer size: 8
}
