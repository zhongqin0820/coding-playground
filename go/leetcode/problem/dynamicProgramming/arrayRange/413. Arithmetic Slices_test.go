package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/arithmetic-slices/description/

// A sequence of number is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.

// For example, these are arithmetic sequence:
// 1, 3, 5, 7, 9
// 7, 7, 7, 7
// 3, -1, -5, -9
// The following sequence is not arithmetic.
// 1, 1, 2, 5, 7

// A zero-indexed array A consisting of N numbers is given. A slice of that array is any pair of integers (P, Q) such that 0 <= P < Q < N.
// A slice (P, Q) of array A is called arithmetic if the sequence:
// A[P], A[p + 1], ..., A[Q - 1], A[Q] is arithmetic. In particular, this means that P + 1 < Q.

// The function should return the number of arithmetic slices in the array A.
func numberOfArithmeticSlices(A []int) int {
	// 在 A[i] - A[i-1] == A[i-1] - A[i-2] 时，dp[i] = dp[i-1] + 1
	if A == nil || len(A) < 3 {
		return 0
	}
	n := len(A)
	dp := make([]int, n)
	for i := 2; i < n; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	// 因为递增子区间不一定以最后一个元素为结尾，可以是任意一个元素结尾
	// 因此需要返回 dp 数组累加的结果
	res := 0
	for _, v := range dp {
		res += v
	}
	return res
}

func numberOfArithmeticSlicesAdv(A []int) int {
	if A == nil || len(A) < 3 {
		return 0
	}
	res, i, j := 0, 0, 0
	for i < len(A) {
		j = i + 2
		for j < len(A) && A[j]-A[j-1] == A[j-1]-A[j-2] {
			j++
		}
		j--
		res += (j - i - 1) * (j - i) / 2
		i = j
	}
	return res
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := []struct {
		A      []int
		output int
	}{
		{[]int{1, 2, 3, 4}, 3},
		// Example:
		// A = [1, 2, 3, 4]
		// return: 3, for 3 arithmetic slices in A: [1, 2, 3], [2, 3, 4] and [1, 2, 3, 4] itself.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numberOfArithmeticSlices(ts.A))
			ast.Equal(ts.output, numberOfArithmeticSlicesAdv(ts.A))
		})
	}
}
