// Slices: How to delete an element from a slice
package main

import (
	"errors"
	"reflect"
	"testing"
)

// DeleteSlice removes an element at a specific index of an slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteSlice(s []int, i int) ([]int, error) {
	if len(s) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(s)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(s[:i], s[i+1:]...), nil
}

func TestDeleteInt(t *testing.T) {
	tests := []struct {
		name    string
		s       []int
		i       int
		want    []int
		wantErr bool
	}{
		{
			name:    "nil slice",
			s:       nil,
			i:       0,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty slice",
			s:       []int{},
			i:       0,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "non empty slice. Index larger than len(a)-1",
			s:       []int{1, 2, 3, 4, 5},
			i:       5,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "non empty slice. Index < 0",
			s:       []int{1, 2, 3, 4, 5},
			i:       -1,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "non empty slice",
			s:       []int{1, 2, 3, 4, 5},
			i:       2,
			want:    []int{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteSlice(tt.s, tt.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
