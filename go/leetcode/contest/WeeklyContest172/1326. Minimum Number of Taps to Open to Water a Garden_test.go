package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-172/problems/minimum-number-of-taps-to-open-to-water-a-garden/
// ref: https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/discuss/484235/JavaC%2B%2BPython-Bad-Description
func minTaps(n int, A []int) int {
	// dp[i] means the smallest number of taps required to water under i;
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n + 2
	}
	dp[0] = 0
	// iterate all taps
	for i := 0; i <= n; i++ {
		for j := helper1326Max(i-A[i]+1, 0); j <= helper1326Min(i+A[i], n); j++ {
			dp[j] = helper1326Min(dp[j], dp[helper1326Max(0, i-A[i])]+1)
		}
	}

	if dp[n] < n+2 {
		return dp[n]
	}
	return -1
}

func helper1326Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func helper1326Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinTaps(t *testing.T) {
	//     Constraints:
	//          1 <= n <= 10^4
	//          ranges.length == n + 1
	//          0 <= ranges[i] <= 100
	tests := []struct {
		n      int
		ranges []int
		output int
	}{
		{5, []int{3, 4, 1, 1, 0, 0}, 1},
		{3, []int{0, 0, 0, 0}, -1},
		{7, []int{1, 2, 1, 0, 2, 1, 0, 1}, 3},
		{8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4}, 2},
		{8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minTaps(ts.n, ts.ranges))
		})
	}
}
