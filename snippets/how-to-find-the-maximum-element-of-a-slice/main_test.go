// Slices: How to find the maximum element of a slice
package main

import (
	"errors"
	"testing"
)

// MaxSlice return the maximum element of slice s
// or an error in case the s is nil or empty
func MaxSlice(s []int) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("cannot find the maximum of a nil or empty slice")
	}

	max := s[0]
	for k := range s {
		if s[k] > max {
			max = s[k]
		}
	}

	return max, nil
}

func TestMaxSlice(t *testing.T) {
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
			s:       []int{1, 4, 5, -34, 2, 100},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxSlice(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxSlice(%#v) error = %v, wantErr %v", tt.s, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxSlice(%#v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
