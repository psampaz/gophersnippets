// Testing: How to test a function that panics
package main

import "testing"

func f(shouldPanic bool) string {
	if shouldPanic {
		panic("function panicked")
	}
	return "function didn't panic"
}

func Test_f(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		f(true)
	})

	t.Run("does not panic", func(t *testing.T) {
		shouldPanic := false
		want := "function didn't panic"
		if got := f(shouldPanic); got != want {
			t.Errorf("f(%v) = %v, want %v", shouldPanic, got, want)
		}
	})

}
