package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minDistanceII(from string, to string) int {
	m := len(from)
	n := len(to)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// dp[i][j] 代表了从 from[:i] → to[:j] 所需要的最小步骤
			// 按照题目给出的 3 种操作方法，分别讨论：
			// 1. 先 *删除* from[:i] 最后的字母，得到 from[:i-1]。
			//    再 from[:i-1] → to[:j] 此方法所需的步骤是
			//    1 + dp[i-1][j]
			// 2. 先 from[:i] → to[:j-1]，
			//    再 *添加* to[j-1] 到 to 的末尾，此方法所需的步骤是
			//    1 + dp[i][j-1]
			// 3. 先 from[i-1] → to[j-1]
			//     3.1 如果 from[i-1] = to[i-1] 的话
			//         无需 *替换*，
			//         总的步骤是 dp[i-1][j-1]
			//     3.2 如果 from[i-1] != to[i-1] 的话
			//         执行 *替换* 操作，把 from[i-1] 替换成 to[j-1]
			//         总的步骤是 1 + dp[i-1][j-1]
			dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])

			replace := 1
			if from[i-1] == to[j-1] {
				replace = 0
			}

			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+replace)
		}
	}

	return dp[m][n]
}

func TestMinDistanceII(t *testing.T) {
	tests := []struct {
		word1  string
		word2  string
		output int
	}{
		{"horse", "ros", 3},
		// Explanation:
		// horse -> rorse (replace 'h' with 'r')
		// rorse -> rose (remove 'r')
		// rose -> ros (remove 'e')
		{"intention", "execution", 5},
		// Explanation:
		// intention -> inention (remove 't')
		// inention -> enention (replace 'i' with 'e')
		// enention -> exention (replace 'n' with 'x')
		// exention -> exection (replace 'n' with 'c')
		// exection -> execution (insert 'u')
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minDistanceII(ts.word1, ts.word2))
		})
	}
}
