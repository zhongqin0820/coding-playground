package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-167/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/
// 计算正方形区域内符合和小于阈值的矩形个数
// ref: https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/discuss/451843/Java-PrefixSum-solution
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	prefixSum := make([][]int, m+1)
	for i, _ := range prefixSum {
		prefixSum[i] = make([]int, n+1)
	}
	// 初始化, prefixSum记录(0,0)到(i,j)区域内的和
	// 想要计算(i,j)到(i+k,j+k)区域内和
	// 等价于计算(i+k,j+k) - (i+k, j-1) - (i-1, j+k) + (i-1,j-1)
	for i := 1; i <= m; i++ {
		sum := 0
		for j := 1; j <= n; j++ {
			sum += mat[i-1][j-1]
			prefixSum[i][j] = prefixSum[i-1][j] + sum
		}
	}
	// 从最大步长k开始
	for k := heper1292Min(m, n); k > 0; k-- {
		for i := 1; i+k <= m; i++ {
			for j := 1; j+k <= n; j++ {
				if prefixSum[i+k][j+k]-prefixSum[i-1][j+k]-prefixSum[i+k][j-1]+
					prefixSum[i-1][j-1] <= threshold {
					return k + 1
				}
			}
		}
	}
	return 0
}

func heper1292Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMaxSideLength(t *testing.T) {
	tests := []struct {
		mat       [][]int
		threshold int
		output    int
	}{
		{[][]int{[]int{1, 1, 3, 2, 4, 3, 2}, []int{1, 1, 3, 2, 4, 3, 2}, []int{1, 1, 3, 2, 4, 3, 2}}, 4, 2},
		{[][]int{[]int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2}}, 1, 0},
		{[][]int{[]int{1, 1, 1, 1}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}}, 6, 3},
		{[][]int{[]int{18, 70}, []int{61, 1}, []int{25, 85}, []int{14, 40}, []int{11, 96}, []int{97, 96}, []int{63, 45}}, 40184, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxSideLength(ts.mat, ts.threshold), ts.threshold)
		})
	}
}
