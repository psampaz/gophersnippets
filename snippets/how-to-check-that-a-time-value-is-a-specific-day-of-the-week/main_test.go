// Time: How to check that a time value is a specific day of the week
package main

import (
	"fmt"
	"time"
)

func Example() {
	// The time package offers a time.Weekday() method which returns
	// a time.Weekday value	https://pkg.go.dev/time?tab=doc#Weekday
	now := time.Now()
	if now.Weekday() == time.Tuesday {
		fmt.Println("The day on Go playground is always Tuesday")
	}
	// Output:
	// The day on Go playground is always Tuesday
}
