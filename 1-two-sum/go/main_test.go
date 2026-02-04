package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   6,
			expected: nil,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			target:   5,
			expected: nil,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4},
			target:   -6,
			expected: []int{1, 3},
		},
		{
			name:     "Mixed numbers",
			nums:     []int{1, -2, 3, 4},
			target:   2,
			expected: []int{1, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := twoSum(test.nums, test.target)

			if len(result) != len(test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
				return
			}

			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("Expected %v, got %v", test.expected, result)
					return
				}
			}
		})
	}
}
