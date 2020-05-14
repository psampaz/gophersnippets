// Slices: How to reverse a slice
package main

import (
	"reflect"
	"testing"
)

func ReverseSlice(s []int) []int {
	if len(s) < 2 {
		return s
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "nil",
			s:    nil,
			want: nil,
		},
		{
			name: "empty",
			s:    []int{},
			want: []int{},
		},
		{
			name: "one element",
			s:    []int{1},
			want: []int{1},
		},
		{
			name: "odd elements",
			s:    []int{1, 2, 3, 4},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "even elements",
			s:    []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSlice(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
