package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainFunction(t *testing.T) {
	t.Run("Test main to return string values of numeronyms(accessibility, Kubernetes, abc) with proper formatting", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf

		main()

		expected := strconv.Quote("[a11y K8s abc]")
		actual := strconv.Quote(buf.String())

		if expected != actual {
			t.Errorf("Unexpected output, expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestNumeronymsFunction(t *testing.T) {
	testCases := []struct {
		description string
		input       []string
		expected    []string
	}{
		{"Test numeronyms with one string input", []string{"accessibility"}, []string{"a11y"}},
		{"Test numeronyms with multiple string input", []string{"accessibility", "Kubernetes", "abc"}, []string{"a11y", "K8s", "abc"}},
		{"Test numeronyms with nil to return nil", nil, nil},
		{"Test numeronyms with empty to return empty string", []string{""}, []string{""}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := numeronyms(testCase.input...)

			for i, actual := range result {
				if testCase.expected[i] != actual {
					t.Errorf("Unexpected output, expected: %s, actual: %s", testCase.expected[i], actual)
				}
			}
		})
	}
}

func TestConvertStringFunction(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    string
	}{
		{"Test convertString with three letters not converted", "abc", "abc"},
		{"Test convertString with two letters not converted", "ab", "ab"},
		{"Test convertString with one letter not converted", "a", "a"},
		{"Test convertString with four letters", "abcd", "a2d"},
		{"Test convertString with numeric values", "1234567", "157"},
		{"Test convertString with non-alphanumeric values", "*!+;'", "*3'"},
		{"Test convertString with empty to return empty string", "", ""},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := convertString(testCase.input)

			if testCase.expected != actual {
				t.Errorf("Unexpected output, expected: %s, actual: %s", testCase.expected, actual)
			}
		})
	}
}
