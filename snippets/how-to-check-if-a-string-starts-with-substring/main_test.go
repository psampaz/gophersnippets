// Strings: How to check if a string starts with substring
package main

import (
	"fmt"
	"strings"
)

func Example() {
	fmt.Println(strings.HasPrefix("abcdefg", "ab"))
	fmt.Println(strings.HasPrefix("abcdefg", "Ab"))
	// Output:
	// true
	// false
}
