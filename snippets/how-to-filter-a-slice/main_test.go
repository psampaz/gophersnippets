// Slices: How to filter a slice
package main

import (
	"reflect"
	"testing"
)

// FilterSlice filters a slice based on a predicate
func FilterSlice(a []int, keep func(x int) bool) []int {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

func TestFilterInt(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		keep func(x int) bool
		want []int
	}{
		{
			name: "nil slice",
			a:    nil,
			keep: func(int) bool { panic("not implemented") },
			want: nil,
		},
		{
			name: "empty slice",
			a:    []int{},
			keep: func(int) bool { panic("not implemented") },
			want: []int{},
		},
		{
			name: "non empty slice",
			a:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			keep: func(x int) bool {
				return x < 5
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterSlice(tt.a, tt.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
