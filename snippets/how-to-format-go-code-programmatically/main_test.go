// Other: How to format Go code programmatically
//
// Code can be formatted programmatically in the same way like running go fmt,
// using the go/format package
package main

import (
	"fmt"
	"go/format"
	"log"
)

func Example() {
	unformatted := `
package main
       import "fmt"

func  main(   )  {
    x :=    12
fmt.Printf(   "%d",   x  )
	}


`
	formatted, err := format.Source([]byte(unformatted))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(formatted))
	// Output:
	// package main
	//
	// import "fmt"
	//
	// func main() {
	//	x := 12
	//	fmt.Printf("%d", x)
	// }
}
