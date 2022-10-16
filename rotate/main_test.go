package binarySearch

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func rotate(nums []int, k int) []int {
	k = k % len(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		logrus.Infof("I: %d", i)
		logrus.Infof("COUNT: %d", count)
		current := i % len(nums)
		logrus.Infof("CURRENT: %d", current)
		prev := nums[current]
		logrus.Infof("PREV: %d", prev)
		for i != current {
			next := (current + k) % len(nums)
			logrus.Infof("NEXT: %d", next)
			temp := nums[next]
			logrus.Infof("nums[next]: %d", nums[next])
			nums[next] = prev
			prev = temp
			current = next
			count++
		}
	}

	return nums
}

func TestRotate(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			/*

			 */
			name:     "4",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{5, 6, 7, 1, 2, 3, 4},
			k:        3,
		},
		{
			/*
				1
				temp = 3
				-1, 99, -100, 3
				2
				temp = -100
				3, 99, -1, -100
			*/
			name:     "1",
			input:    []int{-1, -100, 3, 99},
			expected: []int{3, 99, -1, -100},
			k:        2,
		},
	}

	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.expected, rotate(cc.input, cc.k))
		})
	}
}
