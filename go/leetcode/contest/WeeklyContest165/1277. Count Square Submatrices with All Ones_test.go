package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-square-submatrices-with-all-ones/
// ref: https://leetcode.com/problems/count-square-submatrices-with-all-ones/discuss/441306/Python-DP-solution
// dp[i][j]表示以matrix[i][j]为右下角的最大正方形数
// dp[i][j] = max(dp[i-1][j-1],dp[i-1][j],dp[i][j-1]), A[i][j] == 0
// dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1) + (A[i-1][j-1] == 1)
func countSquares(A [][]int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] > 0 && i > 0 && j > 0 {
				// dp[i][j] = min(dp[i-1][j],dp[i][j-1])
				l := helper1277Min(A[i-1][j], A[i][j-1])
				// dp[i][j] += (A[i-1][j-1] == 1)
				A[i][j] = l + helper1277A(A[i-l][j-l])
			}
			res += A[i][j]
		}
	}
	return res
}

func helper1277Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func helper1277A(a int) int {
	if a > 0 {
		return 1
	}
	return 0
}

func TestCountSquares(t *testing.T) {
	tests := []struct {
		matrix [][]int
		output int
	}{
		{[][]int{[]int{0, 1, 1, 1}, []int{1, 1, 1, 1}, []int{0, 1, 1, 1}}, 15},
		{[][]int{[]int{1, 0, 1}, []int{1, 1, 0}, []int{1, 1, 0}}, 7},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, countSquares(ts.matrix))
		})
	}
}
