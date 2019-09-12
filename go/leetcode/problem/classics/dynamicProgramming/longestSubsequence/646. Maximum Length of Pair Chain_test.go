package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-length-of-pair-chain/description/

// You are given n pairs of numbers. In every pair, the first number is always smaller than the second number.
// Now, we define a pair (c, d) can follow another pair (a, b) if and only if b < c. Chain of pairs can be formed in this fashion.
// Given a set of pairs, find the length longest chain which can be formed. You needn't use up all the given pairs. You can select pairs in any order.

// Note:
// The number of given pairs will be in the range [1, 1000].

func findLongestChain(pairs [][]int) int {
	// 关键是排序！
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] { //第2个数相同，按第一个数降序
			return pairs[i][0] > pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1] //否则按第二个数升序
	})

	res := 0
	b := -1 << 32 //begin
	for i := 0; i < len(pairs); i++ {
		c := pairs[i][0]
		if b < c { //当前的第一个节点可以连上
			res++
			b = pairs[i][1]
		}
	}
	return res
}

func TestFindLongestChain(t *testing.T) {
	tests := []struct {
		pairs  [][]int
		output int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, 2},
		{[][]int{{1, 2}, {2, 2}, {3, 4}}, 2},
		// Explanation: The longest chain is [1,2] -> [3,4]
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findLongestChain(ts.pairs))
		})
	}
}
