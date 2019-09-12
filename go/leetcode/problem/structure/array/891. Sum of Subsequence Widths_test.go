package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-subsequence-widths/

// Given an array of integers A, consider all non-empty subsequences of A.
// For any sequence S, let the width of S be the difference between the maximum and minimum element of S.
// Return the sum of the widths of all subsequences of A.
// As the answer may be very large, return the answer modulo 10^9 + 7.

func sumSubseqWidths(A []int) int {
	mod := int(1e9 + 7)
	sort.Ints(A)

	n := len(A)
	res, bases := 0, 1
	for i := 0; i < n; i++ {
		res += (A[i] - A[n-i-1]) * bases
		res %= mod
		bases = (bases << 1) % mod
	}
	return res
}

// https://leetcode.com/problems/sum-of-subsequence-widths/discuss/161267/C++Java1-line-Python-Sort-and-One-Pass
// For A[i]:
// There are i smaller numbers,
// so there are 2 ^ i sequences in which A[i] is maximum.
// we should do res += A[i] * (2 ^ i)
// There are n - i - 1 bigger numbers,
// so there are 2 ^ (n - i - 1) sequences in which A[i] is minimum.
// we should do res -= A[i] * 2 ^ (n - i - 1)

func TestSumSubseqWidths(t *testing.T) {
	tests := []struct {
		A      []int
		output int
	}{
		{[]int{2, 1, 3}, 6},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sumSubseqWidths(ts.A))
		})
	}
}
