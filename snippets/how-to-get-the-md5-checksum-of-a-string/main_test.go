// Strings: How to get the md5 checksum of a string
package main

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func md5Checksum(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func Test_md5Checksum(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty string",
			s:    "",
			want: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name: "hello world",
			s:    "hello world",
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5Checksum(tt.s); got != tt.want {
				t.Errorf("md5Checksum(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
