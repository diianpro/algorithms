package binarySearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func search(nums []int, target int) int {
	middle := (len(nums) - 1) / 2
	for middle > 0 && middle <= len(nums)-1 {

		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			middle = len(nums[:middle]) / 2
		} else {
			middle += len(nums[middle:]) / 2
		}
	}
	return -1
}

func TestSearch(t *testing.T) {
	cases := []struct {
		name             string
		input            []int
		target, expected int
	}{
		{
			name:     "1",
			input:    []int{-1, 0, 3, 5, 9, 12},
			expected: 4,
		},
	}

	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			shortestSubStrLen := search(cc.input, cc.expected)
			require.Equal(t, cc.expected, shortestSubStrLen)
		})
	}
}
