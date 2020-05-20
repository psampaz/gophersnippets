// Slices: How to split a slice in chunks
package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func SplitSliceInChunks(a []int, chuckSize int) ([][]int, error) {
	if chuckSize < 1 {
		return nil, errors.New("chuckSize must be greater that zero")
	}
	chunks := make([][]int, 0, (len(a)+chuckSize-1)/chuckSize)

	for chuckSize < len(a) {
		a, chunks = a[chuckSize:], append(chunks, a[0:chuckSize:chuckSize])
	}
	chunks = append(chunks, a)
	return chunks, nil
}

func TestSplitSliceInChunks(t *testing.T) {
	tests := []struct {
		name      string
		a         []int
		chuckSize int
		want      [][]int
		wantErr   bool
	}{
		{
			name:      "",
			a:         []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			chuckSize: 1,
			want: [][]int{
				{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9},
			},
			wantErr: false,
		},
		{
			name:      "",
			a:         []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			chuckSize: 10,
			want: [][]int{
				{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			wantErr: false,
		},
		{
			name:      "",
			a:         []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			chuckSize: 3,
			want: [][]int{
				{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9},
			},
			wantErr: false,
		},
		{
			name:      "",
			a:         []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			chuckSize: 5,
			want: [][]int{
				{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitSliceInChunks(tt.a, tt.chuckSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitSliceInChunks(%v, %d) error = %v, wantErr %v", tt.a, tt.chuckSize, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitSliceInChunks(%v, %d) got = %v, want %v", tt.a, tt.chuckSize, got, tt.want)
			}
		})
	}
}

func ExampleSplitSliceInChunks() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	chunks, err := SplitSliceInChunks(a, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", chunks)
	// Output:
	// [[0 1 2 3] [4 5 6 7] [8 9]]
}
