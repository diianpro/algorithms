package sumOddLengthSubarrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func swap(i, j int, arr []int) {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
}

func moveZeroes(arr []int) {
	for lastZeroIndex, cur := 0, 0; cur < len(arr); cur++ {
		if arr[cur] != 0 {
			swap(lastZeroIndex, cur, arr)
			lastZeroIndex++
		}
	}
}

func TestSumOddLengthSubarrays(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		result []int
	}{
		{
			name:   "1",
			arr:    []int{0, 1, 0, 3, 12},
			result: []int{1, 3, 12, 0, 0},
		},
		{
			name:   "2",
			arr:    []int{0},
			result: []int{0},
		},
	}
	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			moveZeroes(cc.arr)
			require.Equal(t, cc.result, cc.arr)
		})
	}
}
