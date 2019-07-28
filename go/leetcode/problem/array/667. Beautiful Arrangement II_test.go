package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/beautiful-arrangement-ii/description/

// Given two integers n and k, you need to construct a list which contains n different positive integers ranging from 1 to n and obeys the following requirement:
// Suppose this list is [a1, a2, a3, ... , an], then the list [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] has exactly k distinct integers.
// If there are multiple answers, print any of them.

func constructArray(n int, k int) []int {
	res := make([]int, 0, n)
	i, j := 1, n
	for i <= j {
		if k%2 == 1 {
			res = append(res, i)
			i++
		} else {
			res = append(res, j)
			j--
		}
		if k > 1 {
			k--
		}
	}

	return res
}

func TestConstructArray(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		output []int
	}{
		{3, 1, []int{1, 2, 3}},
		// Explanation: The [1, 2, 3] has three different positive integers ranging from 1 to 3, and the [1, 1] has exactly 1 distinct integer: 1.
		{3, 2, []int{3, 1, 2}},
		// Explanation: The [1, 3, 2] has three different positive integers ranging from 1 to 3, and the [2, 1] has exactly 2 distinct integers: 1 and 2.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, constructArray(ts.n, ts.k))
		})
	}
}
