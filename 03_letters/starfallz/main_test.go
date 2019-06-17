package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainFunction(t *testing.T) {
	t.Run("Test main to return string values of letters(aba) with proper formatting", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf

		main()

		expected := strconv.Quote(`a:2
b:1`)
		actual := strconv.Quote(buf.String())

		if expected != actual {
			t.Errorf("Unexpected output, expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestLettersFunction(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    map[rune]int
	}{
		{"Test letters to return sorted map of letters based on rune # with count", "abaabazzzzzz", map[rune]int{97: 4, 98: 2, 122: 6}},
		{"Test letters to handle uppercase letters", "aAAaA", map[rune]int{65: 3, 97: 2}},
		{"Test letters to handle non-alphabetical characters", "    !!!*", map[rune]int{32: 4, 33: 3, 42: 1}},
		{"Test letters to return nil on empty string input", "", nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := letters(testCase.input)

			for i, actual := range result {
				if testCase.expected[i] != actual {
					t.Errorf("Unexpected output, expected: %d, actual: %d", testCase.expected[i], actual)
				}
			}
		})
	}
}
