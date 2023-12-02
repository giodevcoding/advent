package helpers

import (
	"fmt"
	"testing"
)

func TestPrefixExistsInList(t *testing.T) {
	var tests = []struct {
		sub      string
		list     []string
		expected bool
	}{
		{"on", []string{"one", "two", "three"}, true},
		{"ree", []string{"one", "two", "three"}, false},
		{"fi", []string{"four", "five", "seven"}, true},
		{"tee", []string{"one", "two", "three"}, false},
		{"three", []string{"one", "two", "three"}, false},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s=%t", test.sub, test.expected)
		t.Run(testname, func(t *testing.T) {
			result := PrefixExistsInList(test.sub, test.list)
			if result != test.expected {
				t.Errorf("GetDigit(%s, []string) = %t; expected %t", test.sub, result, test.expected)
			}
		})
	}
}
