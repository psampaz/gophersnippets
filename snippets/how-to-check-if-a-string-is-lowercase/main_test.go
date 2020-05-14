// Strings: How to check if a string is lowercase
package main

import (
	"testing"
	"unicode"
)

func IsLowercase(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func Test_IsLowercase(t *testing.T) {
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
			want: true,
		},
		{
			name: "one uppercase",
			s:    "abcD",
			want: false,
		},
		{
			name: "all uppercase",
			s:    "ABCDᾺ",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLowercase(tt.s); got != tt.want {
				t.Errorf("IsLowercase(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
