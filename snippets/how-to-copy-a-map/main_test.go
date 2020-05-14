// Maps: How to copy a map
package main

import (
	"reflect"
	"testing"
)

func CopyMap(src map[string]string) map[string]string {
	if src == nil {
		return nil
	}
	// Initialize a new map
	dst := make(map[string]string)
	// range through the source map
	// and copy each element to the new one
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func TestCopyMap(t *testing.T) {
	tests := []struct {
		name string
		src  map[string]string
		want map[string]string
	}{
		{
			name: "nil map",
			src:  nil,
			want: nil,
		},
		{
			name: "empty map",
			src:  make(map[string]string),
			want: make(map[string]string),
		},
		{
			name: "non empty map",
			src: map[string]string{
				"a": "a",
				"b": "b",
			},
			want: map[string]string{
				"a": "a",
				"b": "b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyMap(tt.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyMap(%v) = %v, want %v", tt.src, got, tt.want)
			}
		})
	}
}
