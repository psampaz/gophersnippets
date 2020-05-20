// Strings: How to define a custom string representation for a type
//
// To control the string representation of a type
// the type needs to implement the Stringer interface
// https://pkg.go.dev/fmt?tab=doc#Stringer
// type Stringer interface {
//	String() string
// }
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
package main

import "fmt"

// UserA does not implement the Stringer interface
type UserA struct {
	FirstName string
	LastName  string
}

// UserB implements the Stringer interface
type UserB struct {
	FirstName string
	LastName  string
}

func (u UserB) String() string {
	return fmt.Sprintf("First name is %s, last name is %s", u.FirstName, u.LastName)
}

// BoolA does implements the Stringer interface
type BoolA bool

// BoolB implements the Stringer interface
type BoolB bool

func (b BoolB) String() string {
	if b {
		return "Yes"
	} else {
		return "No"
	}
}

func Example() {
	a := UserA{"John", "Doe"}
	b := UserB{"John", "Doe"}
	c := BoolA(true)
	d := BoolB(true)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// Output:
	// {John Doe}
	// First name is John, last name is Doe
	// true
	// Yes
}
