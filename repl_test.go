package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"HELLO", "WORLD"},
		},
		{
			input:    "hello,world",
			expected: []string{"hello,world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "       ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Expected: %v\n Actual: %v", c.expected, actual)
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Expected: %v\n Actual: %v\n Expected Slice: %v\n Actual Slice: %v", expectedWord, word, c.expected, actual)
				return
			}
		}
	}

}