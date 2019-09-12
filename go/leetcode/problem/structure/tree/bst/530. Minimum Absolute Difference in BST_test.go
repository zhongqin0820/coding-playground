package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

// Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt32, -1

	var in func(*TreeNode)
	in = func(node *TreeNode) {
		if node == nil {
			return
		}
		in(node.Left)
		if pre >= 0 && node.Val-pre < res {
			res = node.Val - pre
		}
		pre = node.Val
		in(node.Right)
	}

	in(root)
	return res
}

func TestGetMinimumDifference(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{1, null, 3, 2}), 1},
		// Explanation:
		// The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
		{Ints2Tree([]int{}), math.MaxInt32},
		{Ints2Tree([]int{1}), math.MaxInt32},
		{Ints2Tree([]int{236, 104, 701, null, 227, null, 911}), 9},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, getMinimumDifference(ts.root))
		})
	}
}
