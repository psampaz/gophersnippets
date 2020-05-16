// Strings: How to count words in a string
package main

import (
	"fmt"
	"strings"
)

// strings.Fields() splits a string around one or more
// consecutive whitespaces (https://golang.org/pkg/unicode/#IsSpace)
// The number of fields is the actual number of words
func CountWords(s string) int {
	return len(strings.Fields(s))
}

// If the words are not separated by whitespaces
// but with comma (a csv row for example) or whatever else,
// you can use the strings.FieldsFunc() and explicitly define
// the points of split
func CountWordsFunc(s string, f func(rune) bool) int {
	return len(strings.FieldsFunc(s, f))
}

func Example_CountWords() {
	wc := CountWords("  How  to count words? \n")
	fmt.Println(wc)
	// Output:
	// 4
}

func Example_CountWordsFunc() {
	f := func(c rune) bool {
		return c == ','
	}
	wc := CountWordsFunc("word1, word2, word3", f)
	fmt.Println(wc)
	// Output:
	// 3
}
