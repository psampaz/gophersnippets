// Strings: How to get the sha256 checksum of a string
package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Sha256Checksum(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func Test_Sha256Checksum(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty string",
			s:    "",
			want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name: "hello world",
			s:    "hello world",
			want: "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha256Checksum(tt.s); got != tt.want {
				t.Errorf("Sha256Checksum(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
