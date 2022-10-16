package binarySearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	middle := (left + right) / 2
	for left <= right {
		middle = (left + right) / 2

		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return right + 1
}

func TestSearch(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{
			name:     "4",
			input:    []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "1",
			input:    []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "2",
			input:    []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "1",
			input:    []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
	}

	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.expected, searchInsert(cc.input, cc.target))
		})
	}
}
