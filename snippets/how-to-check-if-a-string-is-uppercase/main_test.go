// Strings: How to check if a string is uppercase
package main

import (
	"testing"
	"unicode"
)

func IsUppercase(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

func Test_IsUppercase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty",
			s:    "",
			want: false,
		},
		{
			name: "all lowercase",
			s:    "abc",
			want: false,
		},
		{
			name: "one lowecase",
			s:    "ABCd",
			want: false,
		},
		{
			name: "all uppercase",
			s:    "ABCD",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUppercase(tt.s); got != tt.want {
				t.Errorf("IsUppercase(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
