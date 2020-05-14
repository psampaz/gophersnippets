// Interfaces: How to check if a type satisfies an interface at runtime
package main

import (
	"fmt"
)

type Checker interface {
	Check() bool
}

type P struct{}

func (p *P) Check() bool {
	return true
}

type V struct{}

func (v V) Check() bool {
	return true
}

type N struct{}

// References:
// https://golang.org/ref/spec#Method_sets
// https://golang.org/ref/spec#Type_assertions
// https://www.ardanlabs.com/blog/2017/07/interface-semantics.html
func Example() {
	// According to the Go language specification:
	// A type may have a method set associated with it.
	// The method set of an interface type is its interface.
	// The method set of any other type T consists of all methods
	// declared with receiver type T.
	// The method set of the corresponding pointer type *T is 
	// the set of all methods declared
	// with receiver *T or T (that is, it also contains the method set of T).
	//
	// In other words:
	// A type T satisfies an interface only if the interface methods are 
	// implemented using value receiver.
	// A type *T satisfies an interface only if the interface methods are
	// implemented using value or pointer receiver

	// P implements the Checker interface using pointer receiver (*P)
	// so *P satisfies the checker interface
	// we assign the value &P to an empty interface
	var pp interface{} = &P{}
	// we use type assertion to check if the value of the interface is type Checker
	if _, ok := pp.(Checker); ok {
		fmt.Println("*P implements the checker interface")
	} else {
		fmt.Println("*P does not implement the checker interface")
	}

	// P implements the Checker interface using pointer receiver (*P)
	// so P does not satisfy the checker interface
	var pv interface{} = P{}

	if _, ok := pv.(Checker); ok {
		fmt.Println("P implements the checker interface")
	} else {
		fmt.Println("P does not implement the checker interface")
	}

	// V implements the Checker interface using value receiver
	// so *V satisfies the interface
	var vp interface{} = &V{}

	if _, ok := vp.(Checker); ok {
		fmt.Println("*V implements the checker interface")
	} else {
		fmt.Println("*V does not implement the checker interface")
	}

	// V implements the Checker interface using value receiver
	// so V satisfies the interface
	var vv interface{} = V{}

	if _, ok := vv.(Checker); ok {
		fmt.Println("V implements the checker interface")
	} else {
		fmt.Println("V does not implement the checker interface")
	}

	// N does not implement the Checker interface at all
	// so *N does not satisfy the interface
	var np interface{} = &N{}

	if _, ok := np.(Checker); ok {
		fmt.Println("*N implements the checker interface")
	} else {
		fmt.Println("*N does not implement the checker interface")
	}

	// N does not implement the Checker interface at all
	// so N does not satisfy the interface
	var nv interface{} = N{}

	if _, ok := nv.(Checker); ok {
		fmt.Println("N implements the checker interface")
	} else {
		fmt.Println("N does not implement the checker interface")
	}
	// Output:
	// *P implements the checker interface
	// P does not implement the checker interface
	// *V implements the checker interface
	// V implements the checker interface
	// *N does not implement the checker interface
	// N does not implement the checker interface
}
