// Conversions: How to convert int to string
package main

import (
	"fmt"
	"strconv"
)

func Example_itoa() {
	// Converting an int to a string can be done
	// using the `strconv.Itoa` (integer to ascii) function
	a := strconv.Itoa(1234)
	fmt.Printf("%q\n", a)
	// Output:
	// "1234"
}

func Example_formatInt() {
	// `strconv.Itoa` is calls internally
	// `FormatInt(int64(i), 10)` so another way
	//  to  convert an int to  a string is the following:
	b := strconv.FormatInt(int64(1234), 10)
	fmt.Printf("%q\n", b)
	// Output:
	// "1234"
}

func Example_sprintf() {
	// A third way would be to use the `fmt.Sprintf` method:
	c := fmt.Sprintf("%d", 1234)
	fmt.Printf("%q\n", c)
	// Output:
	// "1234"
}
