package sumOddLengthSubarrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	for i := 1; i <= len(arr); i += 2 {
		start := 0
		end := start + i
		for end <= len(arr) {
			for i := start; i < end; i++ {
				sum += arr[i]
			}
			start++
			end++
		}
	}
	return sum
}

func TestSumOddLengthSubarrays(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		result int
	}{
		{
			name:   "1",
			arr:    []int{1, 4, 2, 5, 3},
			result: 58,
		},
		{
			name:   "2",
			arr:    []int{1, 2},
			result: 3,
		},
	}
	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.result, sumOddLengthSubarrays(cc.arr))
		})
	}
}
