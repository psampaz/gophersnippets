// Slices: How to remove duplicate elements from a slice
package main

import (
	"reflect"
	"testing"
)

func DeduplicateSlice(s []int) []int {
	if len(s) < 2 {
		return s
	}

	seen := make(map[int]struct{})

	j := 0
	for k := range s {
		if _, ok := seen[s[k]]; ok {
			continue
		}
		seen[s[k]] = struct{}{}
		s[j] = s[k]
		j++
	}

	return s[:j]
}

func TestDeduplicateSlice(t *testing.T) {
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
			name: "multiple elements",
			s:    []int{1, 2, 3, 3, 2, 5, 3},
			want: []int{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateSlice(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
