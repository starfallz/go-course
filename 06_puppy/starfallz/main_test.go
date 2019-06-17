package main

import (
	"testing"
)

func TestCreatePuppyFunction(t *testing.T) {
	t.Run("Test Creation of same Puppy twice", func(t *testing.T) {
		testCases := []struct {
			description string
			input       Puppy
			expected    string
		}{
			{"Test Creation of Puppy", Puppy{111, "Pug", "Black", 3000.00}, "<nil>"},
			{"Test Creation of same Puppy twice", Puppy{111, "Pug", "Black", 3000.00}, "ERROR: Create failed, puppy with ID 111 already exist"},
		}

		s := &MapStore{}

		for _, testCase := range testCases {
			actual := s.CreatePuppy(&testCase.input)

			if testCase.expected != actual.Error() {
				t.Errorf("Unexpected output, expected: %s, actual: %s", testCase.expected, actual.Error())
			}
		}
	})
}
