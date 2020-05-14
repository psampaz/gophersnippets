// Strings: How to get the sha1 checksum of a string
package main

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func Sha1Checksum(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func Test_Sha1Checksum(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty string",
			s:    "",
			want: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		{
			name: "hello world",
			s:    "hello world",
			want: "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha1Checksum(tt.s); got != tt.want {
				t.Errorf("Sha1Checksum(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
