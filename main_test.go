package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"Hello World", "dlroW olleH", false},
		{"123456", "654321", false},
		{"", "", true},
		{"  ", "", true},
	}

	for _, test := range tests {
		actual, err := reverse(test.input)

		if test.err == true && err == nil {
			t.Errorf("we expect an error with input = %q", test.input)
		}

		if test.err == false && err != nil {
			t.Errorf("we do not expect an error with input = %q", test.input)
		}

		if actual != test.expected {
			t.Errorf("reverse(%q) = %q but expected %q", test.input, actual, test.expected)
		}
	}
}
