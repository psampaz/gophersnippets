// Strings: How to pretty print JSON
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     string
	Lessons []string
}

func Example() {
	s := Student{
		Name: "John",
		Age:  "17",
		Lessons: []string{
			"Mathematics",
			"Computer science",
			"Philoshopy",
		},
	}

	// By using json.Marshal the output will be one line json string
	// which is difficult to read or debug
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nUgly print:\n%s\n", jsonBytes)

	// The easiest way to achieve a human readable and pretty print
	// is to use the json.MarshalIndent
	jsonBytes, err = json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nPretty print:\n%s\n", jsonBytes)
	// Output:
	//
	// Ugly print:
	// {"Name":"John","Age":"17","Lessons":["Mathematics","Computer science","Philoshopy"]}
	//
	// Pretty print:
	// {
	//	"Name": "John",
	//	"Age": "17",
	//	"Lessons": [
	//		"Mathematics",
	//		"Computer science",
	//		"Philoshopy"
	//	]
	// }
}
