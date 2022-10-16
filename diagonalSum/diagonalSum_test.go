package sumOddLengthSubarrays

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func diagonalSum(mat [][]int) int {
	sum := 0
	for i, j := 0, len(mat)-1; i < len(mat); i, j = i+1, j-1 {
		if i == j {
			sum += mat[i][i]
		} else {
			sum += mat[i][i] + mat[i][j]
		}
	}
	return sum
}
func TestName(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	sum := 0
	for i, j := 0, len(mat)-1; i < len(mat); i, j = i+1, j-1 {
		if i == j {
			sum += mat[i][i]
		} else {
			sum += mat[i][i] + mat[i][j]
		}
	}
	fmt.Println("SUM: ", sum)
}

func TestMaximumWealth(t *testing.T) {
	cases := []struct {
		name   string
		mat    [][]int
		result int
	}{
		{
			name: "1",
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			result: 25,
		},
		{
			name: "2",
			mat: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			result: 8,
		},
	}
	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.result, diagonalSum(cc.mat))
		})
	}
}
