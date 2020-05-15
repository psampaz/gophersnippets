// Other: How to print the binary representation of an integer
package main

import (
	"fmt"
)

func Example() {
	// The %b verb print the binary representation of an integer
	for x:=0; x<16;x++ {
		fmt.Printf("%b\n", x)
	}
	// If you want the output to have the same length, 8 bits for example
	// use the %08b notation which means:
	// print binary, use 8 digits for the output, pad with leading zeros
	fmt.Println("fixed length of 8 digits:")
	for x:=0; x<16;x++ {
		fmt.Printf("%08b\n", x)
	}
	// Output:
	// 0
	// 1
	// 10
	// 11
	// 100
	// 101
	// 110
	// 111
	// 1000
	// 1001
	// 1010
	// 1011
	// 1100
	// 1101
	// 1110
	// 1111
	// fixed length of 8 digits:
	// 00000000
	// 00000001
	// 00000010
	// 00000011
	// 00000100
	// 00000101
	// 00000110
	// 00000111
	// 00001000
	// 00001001
	// 00001010
	// 00001011
	// 00001100
	// 00001101
	// 00001110
	// 00001111
}
