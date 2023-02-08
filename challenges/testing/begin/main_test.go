// challenges/testing/begin/main_test.go
package main

import "testing"

type testcase struct {
	input string
	want int
}

func assert(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	cases := []testcase{
		{"#00", 0},
		{"a2_32_^_&/", 1},
		{"812_%6ab/", 2},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			lc := letterCounter{identifier: "letters"}
			got := lc.count(tc.input)
			assert(t, got, tc.want)
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	cases := []testcase{
		{"#00", 2},
		{"abc_1,?", 1},
		{"abc_12&8_^", 3},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			c := numberCounter{designation: "numbers"}
			got := c.count(tc.input)
			assert(t, got, tc.want)
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	cases := []testcase{
		{"#00", 1},
		{"abc_1,?", 3},
		{"abc_12&8_^_", 5},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			c := symbolCounter{label: "symbols"}
			got := c.count(tc.input)
			assert(t, got, tc.want)
		})
	}
}