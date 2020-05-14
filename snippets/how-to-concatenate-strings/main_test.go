// Strings: How to concatenate strings
package main

import (
	"fmt"
	"strings"
)

func Example() {
	// Create a slice of the strings that you want to join together
	strs := []string{"aa", "bb", "cc"}
	//Use the strings.Join function to join them, by passing the desired separator
	fmt.Println(strings.Join(strs, "-"))
	fmt.Println(strings.Join(strs, ", "))
	fmt.Println(strings.Join(strs, ""))
	// Output:
	// aa-bb-cc
	// aa, bb, cc
	// aabbcc
}
