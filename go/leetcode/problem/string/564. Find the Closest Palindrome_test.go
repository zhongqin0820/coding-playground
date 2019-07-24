package problem

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-closest-palindrome/

// Given an integer n, find the closest integer (not including itself), which is a palindrome.
// The 'closest' is defined as absolute difference minimized between two integers.

// Note:
// The input n is a positive integer represented by string, whose length will not exceed 18.
// If there is a tie, return the smaller one as answer.

func nearestPalindromic(s string) string {
	// 获得位数为n-1与位数为n的最值；如n=3时
	// 99 101 999 1001
	n, base := len(s), 1
	for i := 1; i < n; i++ {
		base *= 10
	}
	candidates := make([]int, 4)
	candidates[0], candidates[1] = base-1, base+1
	base *= 10
	candidates[2], candidates[3] = base-1, base+1
	// 所有可能的值
	candidates = append(candidates, helper564(s)...)
	//
	num, _ := strconv.Atoi(s)
	delta := func(x int) int {
		if x > num {
			return x - num
		}
		return num - x
	}

	res := math.MaxInt64
	for _, cand := range candidates {
		if cand == num {
			continue
		}
		if delta(cand) < delta(res) || (delta(cand) == delta(res) && cand < res) {
			res = cand
		}
	}

	return strconv.Itoa(res)
}

func helper564(s string) []int {
	// 获取s附近的值，以前一半为模板，因为前一半在高位
	n, prefix := len(s), s[:(len(s)+1)/2]
	p, _ := strconv.Atoi(prefix)
	candidates, ps := make([]int, 3), make([]int, 3)
	candidates[0], ps[0] = p-1, p-1
	candidates[1], ps[1] = p, p
	candidates[2], ps[2] = p+1, p+1
	// 奇数位
	if n%2 == 1 {
		for i := range ps {
			ps[i] /= 10
		}
	}
	// 回文补充
	for i := range candidates {
		for ps[i] > 0 {
			candidates[i] = candidates[i]*10 + ps[i]%10
			ps[i] /= 10
		}
	}
	return candidates
}

func TestNearestPalindromic(t *testing.T) {
	tests := []struct {
		n      string
		output string
	}{
		{"123", "121"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, nearestPalindromic(ts.n))
		})
	}
}
