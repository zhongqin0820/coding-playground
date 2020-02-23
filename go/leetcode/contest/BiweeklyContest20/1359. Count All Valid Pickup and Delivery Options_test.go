package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/biweekly-contest-20/problems/count-all-valid-pickup-and-delivery-options/
// ref: https://leetcode.com/problems/count-all-valid-pickup-and-delivery-options/discuss/516968/JavaC%2B%2BPython-Straight-Forward
// int countOrders(int n) {
//     int inv = (mod + 1) / 2;
//     int ans = 1;
//     for(int i = 1; i <= 2 * n; ++i) {
//         ans = ans * 1ll * i % mod;
//     }
//     for(int i = 1; i <= n; ++i) {
//         ans = ans * 1ll * inv % mod;
//     }
//     return ans;
// }
var mod = 1000000007

func countOrders(n int) int {
	inv := (mod + 1) / 2
	ans := 1
	for i := 1; i <= 2*n; i++ {
		ans = ans * i % mod
	}
	for i := 1; i <= n; i++ {
		ans = ans * inv % mod
	}
	return ans % mod
}

func TestCountOrders(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{1, 1},
		{2, 6},
		{3, 90},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, countOrders(ts.n))
		})
	}
}
