// Strings: How to reverse a string
package main

import "testing"

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty",
			s:    "",
			want: "",
		},
		{
			name: "non empty ascii",
			s:    "abcd",
			want: "dcba",
		},
		{
			name: "non empty utf8",
			s:    "καλημέρα",
			want: "αρέμηλακ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.s); got != tt.want {
				t.Errorf("ReverseString(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
