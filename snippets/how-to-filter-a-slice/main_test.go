// Slices: How to filter a slice
package main

import (
	"reflect"
	"testing"
)

// FilterSlice filters a slice based on a predicate
func FilterSlice(s []int, keep func(x int) bool) []int {
	if len(s) == 0 {
		return s
	}

	n := 0
	for _, v := range s {
		if keep(v) {
			s[n] = v
			n++
		}
	}

	return s[:n]
}

func TestFilterInt(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		keep func(x int) bool
		want []int
	}{
		{
			name: "nil slice",
			s:    nil,
			keep: func(int) bool { panic("not implemented") },
			want: nil,
		},
		{
			name: "empty slice",
			s:    []int{},
			keep: func(int) bool { panic("not implemented") },
			want: []int{},
		},
		{
			name: "non empty slice",
			s:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			keep: func(x int) bool {
				return x < 5
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterSlice(tt.s, tt.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
