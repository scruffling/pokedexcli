package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		// add more cases here
		//
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			// error and continue to next case
			t.Errorf("Length mismatch, exptected: %v, actual: %v",
				c.expected, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if diff := cmp.Diff(expectedWord, word); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		}
	}
}
