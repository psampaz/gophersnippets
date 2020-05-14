// Slices: How to shuffle a slice
package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func ShuffleSlice(a []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	return a
}

func TestShuffleSlice(t *testing.T) {
	// start with a sorted slice
	s := []int{1, 2, 3, 4, 5}

	// make a copy of the s slice and keep it as reference
	// for later comparison since the ShuffleSlice is an in place operation
	// and it will change s slice
	// See https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice
	original := append(s[:0:0], s...)

	// shuffle the slice and check that it is different from the original
	shuffled := ShuffleSlice(s)
	if reflect.DeepEqual(shuffled, original) {
		t.Errorf("Shuffled slice should be different from the original")
	}

	// Sort the shuffled slice in order to ensure that its elements are the same with the original
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
	}
}
