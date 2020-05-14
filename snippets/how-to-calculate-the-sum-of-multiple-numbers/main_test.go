// Numbers: How to calculate the sum of multiple numbers
package main

import (
	"fmt"
	"testing"
)

// References:
// https://golang.org/ref/spec#Passing_arguments_to_..._parameters
func AddNumbers(s ...int) int {
	sum := 0
	for _, x := range s {
		sum += x
	}
	return sum
}

func TestAddNumbers(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want int
	}{
		{
			name: "nil",
			s:    nil,
			want: 0,
		},
		{
			name: "single number",
			s:    []int{1},
			want: 1,
		},
		{
			name: "multiple numbers",
			s:    []int{4, 3, 2, 1, 0},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddNumbers(tt.s...); got != tt.want {
				t.Errorf("AddNumbers(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func ExampleAddNumbers() {
	fmt.Println(AddNumbers(4, 3, 2, 1, 0))
	fmt.Println(AddNumbers([]int{4, 3, 2, 1, 0}...))
	// Output:
	// 10
	// 10
}
