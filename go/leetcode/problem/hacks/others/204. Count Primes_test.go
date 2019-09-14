package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-primes/description/
func countPrimes(n int) int {
	res, p := 0, make([]bool, n)
	for i := 2; i < n; i++ {
		if !p[i] {
			res++
			// 埃拉托斯特尼筛法在每次找到一个素数时，将能被素数整除的数排除掉
			for j := 2; i*j < n; j++ {
				p[i*j] = true
			}
		}
	}
	return res
}

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{10, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, countPrimes(ts.n))
		})
	}
}
