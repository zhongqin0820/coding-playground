package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-171/problems/minimum-flips-to-make-a-or-b-equal-to-c/
// ref: https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/discuss/477690/Java-Bit-manipulation-w-explanation-and-analysis.
func minFlips(a int, b int, c int) int {
	res, ab, equal := 0, a|b, (a|b)^c
	var i uint32
	for ; i < 31; i++ {
		mask := 1 << i
		if equal&mask > 0 { // need to flip
			if (ab&mask) < (c&mask) || (a&mask) != (b&mask) {
				res += 1
			} else {
				res += 2
			}
			ab ^= mask
		}
	}
	return res
}

func TestMinFlips(t *testing.T) {
	// Constraints:
	//     1 <= a <= 10^9
	//     1 <= b <= 10^9
	//     1 <= c <= 10^9
	tests := []struct {
		a      int
		b      int
		c      int
		output int
	}{
		{2, 6, 5, 3},
		{4, 2, 7, 1},
		{1, 2, 3, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minFlips(ts.a, ts.b, ts.c))
		})
	}
}
