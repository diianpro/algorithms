package sumOddLengthSubarrays

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func matrixReshape(mat [][]int, r int, c int) [][]int {
	res := make([][]int, r)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, c)
	}
	if len(mat) == 0 || r*c != len(mat)*len(mat[0]) {
		return mat
	}
	rows, cols := 0, 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			res[rows][cols] = mat[i][j]
			cols++
			if cols == c {
				rows++
				cols = 0
			}
		}
	}
	return res
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
		r, c   int
		result [][]int
	}{
		{
			name: "1",
			mat: [][]int{
				{1, 2},
				{3, 4},
			},
			r:      1,
			c:      4,
			result: [][]int{{1, 2, 3, 4}},
		},
	}
	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.result, matrixReshape(cc.mat, cc.r, cc.c))
		})
	}
}
