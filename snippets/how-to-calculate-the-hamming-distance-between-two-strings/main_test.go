// Strings: How to calculate the hamming distance between two strings
package main

import (
	"errors"
	"testing"
)

// https://en.wikipedia.org/wiki/Hamming_distance:
// Τhe Hamming distance between two strings of equal length
// is the number of positions at which the corresponding symbols are different.
// In other words, it measures the minimum number of substitutions
// required to change one string into the other
//
// Hamming distance function calculates the rune based
// hamming distance between strings a and b
func HammingDistance(a string, b string) (int, error) {
	// stings in Go are slices of bytes so we have first to convert them to runes
	// read more here https://blog.golang.org/strings
	ra := []rune(a)
	rb := []rune(b)
	if len(ra) != len(rb) {
		return 0, errors.New("strings do not have the same length")
	}

	var distance int
	for i := range ra {
		if rb[i] != ra[i] {
			distance++
		}
	}
	return distance, nil
}

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		name    string
		a       string
		b       string
		want    int
		wantErr bool
	}{
		{
			name:    "no equal length",
			a:       "abc",
			b:       "abcd",
			want:    0,
			wantErr: true,
		},
		{
			name:    "same strings ascii",
			a:       "abcd",
			b:       "abcd",
			want:    0,
			wantErr: false,
		},
		{
			name:    "one character different ascii",
			a:       "Abcd",
			b:       "abcd",
			want:    1,
			wantErr: false,
		},
		{
			name:    "all characters different ascii",
			a:       "ABCD",
			b:       "abcd",
			want:    4,
			wantErr: false,
		},
		{
			name:    "same strings utf8",
			a:       "Καλημέρα",
			b:       "Καλημέρα",
			want:    0,
			wantErr: false,
		},
		{
			name:    "one character different utf8",
			a:       "Καλημέρα",
			b:       "Καλημερα",
			want:    1,
			wantErr: false,
		},
		{
			name:    "all characters different utf8",
			a:       "ΚΑΛΗΜΕΡΑ",
			b:       "καλημέρα",
			want:    8,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HammingDistance(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"HammingDistance(%s %s) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr,
				)
				return
			}
			if got != tt.want {
				t.Errorf("HammingDistance(%s %s) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
