package helpers

import (
	"fmt"
	"testing"
)

func TestIntSlicesEqual(t *testing.T) {
	type (
		Input struct {
			sliceA []int
			sliceB []int
		}

		Test struct {
			input    Input
			expected bool
		}
	)

	var tests = []Test{
		{Input{[]int{1, 2, 3}, []int{1, 2, 3}}, true},
		{Input{[]int{1, 2, 3}, []int{2, 3, 4}}, false},
		{Input{[]int{1, 2, 3, 5, 6}, []int{1, 2, 3, 5, 6}}, true},
		{Input{[]int{1, 2, 3, 5, 6}, []int{1, 2, 3, 4, 5, 6}}, false},
		{Input{[]int{}, []int{}}, true},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("IntSlicesEqual: %v == %v", test.input.sliceA, test.input.sliceB)
		t.Run(testname, func(t *testing.T) {
			result := IntSlicesEqual(test.input.sliceA, test.input.sliceB)
			if result != test.expected {
				t.Errorf("IntSlicesEqual(%v, %v) = %t; expected %t", test.input.sliceA, test.input.sliceB, result, test.expected)
			}
		})
	}
}
