package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"Hello World", "dlroW olleH", false},
		{"123456", "654321", false},
		{"A", "A", false},
		{"", "", true},
		{"  ", "", true},
	}

	for _, test := range tests {
		actual, err := reverse(test.input)
		assert := assert.New(t)

		assert.Equal(test.err == true, err != nil, fmt.Sprintf("we expect an error with input = %q", test.input))
		assert.Equal(test.err == false, err == nil, fmt.Sprintf("we do not expect an error with input = %q", test.input))
		assert.Equal(actual, test.expected, fmt.Sprintf("reverse(%q) = %q but expected %q", test.input, actual, test.expected))
	}
}
