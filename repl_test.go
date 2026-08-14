package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"not trimmed input": {input: " hello world ", expected: []string{"hello", "world"}},
		"simple":            {input: "Hey you there", expected: []string{"hey", "you", "there"}},
		"double space":      {input: "hello  there", expected: []string{"hello", "there"}},
		"caps":              {input: "HELLO there", expected: []string{"hello", "there"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CleanInput(tc.input)
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}
