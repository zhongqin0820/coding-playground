package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-171/problems/minimum-distance-to-type-a-word-using-two-fingers/
// ref1: https://leetcode.com/problems/minimum-distance-to-type-a-word-using-two-fingers/discuss/477659/Top-Down-DP
// ref2: https://leetcode.com/problems/minimum-distance-to-type-a-word-using-two-fingers/discuss/477652/JavaC%2B%2BPython-DP-Solution-O(1)-Space
func minimumDistance(word string) int {
	var dp [26]int
	res, save, n := 0, 0, len(word)
	for i := 0; i < n-1; i++ {
		b, c := int(word[i]-'A'), int(word[i+1]-'A')
		for a := 0; a < 26; a++ {
			if dp[b] < dp[a]+helper1320(b, c)-helper1320(a, c) {
				dp[b] = dp[a] + helper1320(b, c) - helper1320(a, c)
			}
		}
		if save < dp[b] {
			save = dp[b]
		}
		res += helper1320(b, c)
	}
	return res - save
}

func helper1320(a, b int) int {
	n, m := a/6-b/6, a%6-b%6
	if n < 0 {
		n = -n
	}
	if m < 0 {
		m = -m
	}
	return n + m
}

func TestMinimumDistance(t *testing.T) {
	// Constraints:
	//     2 <= word.length <= 300
	//     Each word[i] is an English uppercase letter.
	tests := []struct {
		word   string
		output int
	}{
		{"CAKE", 3},
		{"HAPPY", 6},
		{"NEW", 3},
		{"YEAR", 7},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minimumDistance(ts.word))
		})
	}
}
