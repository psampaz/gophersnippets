// Maps: How to create a set using a map
package main

import "fmt"

// A set is a unordered collection of unique data.
// Go does not have a built-in set data structure
// but you can easily create one using a map
func Example() {
	// The map keys are a unique unordered collection.
	// So if you just ignore the values, a map can be viewed as a set
	// The most space optimized way to do this is to use the empty struct
	// as the map value type, since the empty struct (a struct with no fields)
	// occupies zero bytes of storage
	// https://dave.cheney.net/2014/03/25/the-empty-struct
	set := make(map[string]struct{})
	fmt.Println(set)
	// Add new elements in the set
	set["a"] = struct{}{}
	set["b"] = struct{}{}
	set["c"] = struct{}{}
	fmt.Println(set)

	// Check if an element exists in the set
	if _, exists := set["a"]; exists {
		fmt.Println("a exists in the set")
	}
	// Delete an element from the set
	delete(set, "b")
	fmt.Println(set)
	// Output:
	// map[]
	// map[a:{} b:{} c:{}]
	// a exists in the set
	// map[a:{} c:{}]
}
