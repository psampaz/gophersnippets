// Logging: how to disable log output
package main

import (
	"io/ioutil"
	"log"
)

func Example() {
	// If you want to disable log output (during test for example)
	// you have to set the logger ouput to ioutil.Discard
	// which is an io.Writer on which all Write calls succeed
	// without doing anything
	log.SetOutput(ioutil.Discard)
	log.Println("log disabled")
	// Output:
	//
}
