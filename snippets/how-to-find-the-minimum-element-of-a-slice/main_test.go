// Slices: How to find the minimum element of a slice
package main

import (
	"errors"
	"testing"
)

// MaxSlice return the minimum element of slice s
// or an error in case the s is nil or empty
func MinSlice(s []int) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("cannot find the minimum of a nil or empty slice")
	}

	min := s[0]
	for k := range s {
		if s[k] < min {
			min = s[k]
		}
	}

	return min, nil
}

func TestMinSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       []int
		want    int
		wantErr bool
	}{
		{
			name:    "nil",
			s:       nil,
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty",
			s:       []int{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "non empty slice",
			s:       []int{1, 3, 2, -4, 6, 5},
			want:    -4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinSlice(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinSlice(%#v) error = %v, wantErr %v", tt.s, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxSlice(%#v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
