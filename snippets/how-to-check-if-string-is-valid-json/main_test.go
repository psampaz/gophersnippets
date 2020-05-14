// Strings: How to check if string is valid JSON
package main

import (
	"encoding/json"
	"fmt"
)

func Example() {
	valid := `{"foo":"bar"}`
	invalid := `}{`

	fmt.Println(json.Valid([]byte(valid)))
	fmt.Println(json.Valid([]byte(invalid)))
	// Output:
	// true
	// false
}
