package binarySearch

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func sortedSquares(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	left, rigth := 0, l-1
	for i := l - 1; i >= 0; i-- {
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[rigth])) {
			res[i] = nums[rigth] * nums[rigth]
			rigth--
		} else {
			res[i] = nums[left] * nums[left]
			left++
		}
	}

	return res
}

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "4",
			input:    []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "1",
			input:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.expected, sortedSquares(cc.input))
		})
	}
}
