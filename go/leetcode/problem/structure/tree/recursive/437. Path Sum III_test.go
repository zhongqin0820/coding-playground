package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/path-sum-iii/description/

// You are given a binary tree in which each node contains an integer value.

// Find the number of paths that sum to a given value.

// The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

// The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSumIII(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper437(root, sum) + pathSumIII(root.Left, sum) + pathSumIII(root.Right, sum)
}

func helper437(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res++
	}
	res += helper437(root.Left, sum-root.Val) + helper437(root.Right, sum-root.Val)
	return res
}

func TestPathSumIII(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		sum    int
		output int
	}{
		{Ints2Tree([]int{10, 5, -3, 3, 2, null, 11, 3, -2, null, 1}), 8, 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, pathSumIII(ts.root, ts.sum))
		})
	}
}
