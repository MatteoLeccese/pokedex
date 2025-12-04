package main

import "testing"

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hi Babe",
			expected: []string{"hi", "babe"},
		},
		{
			input:    "                                 hello                                              world                                        .  ",
			expected: []string{"hello", "world", "."},
		},
	}

	for _, testCase := range testCases {
		actual := cleanInput(testCase.input)

		if len(actual) != len(testCase.expected) {
			t.Errorf("For input '%s', expected length %d but got %d", testCase.input, len(testCase.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := testCase.expected[i]
			if word != expectedWord {
				t.Errorf("For input '%s', at index %d, expected '%s' but got '%s'", testCase.input, i, expectedWord, word)
				continue
			}
		}
	}
}
