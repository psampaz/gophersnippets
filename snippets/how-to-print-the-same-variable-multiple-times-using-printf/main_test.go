// Other: How to print the same variable multiple times using printf
package main

import (
	"fmt"
)

func Example() {
	// You can print multiple times the same variable using argument indexes
	// The notation [n] before the verb, means use the nth argument after the format.
	// The above applies to Printf, Sprintf, and Fprintf
	fmt.Printf("%[1]d %[1]d\n", 5)
	fmt.Printf("%[2]d %[2]d %[1]d %[1]d\n", 5, 6)
	// Output:
	// 5 5
	// 6 6 5 5
}
