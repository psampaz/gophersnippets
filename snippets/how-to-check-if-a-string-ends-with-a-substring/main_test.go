// Strings: How to check if a string ends with a substring
package main

import (
	"fmt"
	"strings"
)

func Example() {
	fmt.Println(strings.HasSuffix("abcdefg", "fg"))
	fmt.Println(strings.HasSuffix("abcdefg", "Fg"))
	// Output:
	// true
	// false
}
