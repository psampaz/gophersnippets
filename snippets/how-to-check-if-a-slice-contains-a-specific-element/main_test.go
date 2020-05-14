// Slices: How to check if a slice contains a specific element
package main

import "testing"

func SliceContains(s []int, x int) bool {
	if len(s) == 0 {
		return false
	}
	for k := range s {
		if s[k] == x {
			return true
		}
	}
	return false
}

func TestSliceContains(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		x    int
		want bool
	}{
		{
			name: "nil",
			s:    nil,
			x:    1,
			want: false,
		},
		{
			name: "empty",
			s:    []int{},
			x:    1,
			want: false,
		},
		{
			name: "value exists",
			s:    []int{2, 1},
			x:    1,
			want: true,
		},
		{
			name: "value does not exist",
			s:    []int{2, 1},
			x:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContains(tt.s, tt.x); got != tt.want {
				t.Errorf("SliceContains(%v, %v) = %v, want %v", tt.s, tt.x, got, tt.want)
			}
		})
	}
}
