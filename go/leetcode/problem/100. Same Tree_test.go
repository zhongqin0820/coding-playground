package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// // https://leetcode.com/problems/same-tree/

// Given two binary trees, write a function to check if they are the same or not.

// Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) && p.Val == q.Val

}

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p, q   *TreeNode
		output bool
	}{
		{Ints2Tree([]int{}), Ints2Tree([]int{}), true},
		{Ints2Tree([]int{1, 2, 3}), Ints2Tree([]int{1, 2, 3}), true},
		{Ints2Tree([]int{1, 2}), Ints2Tree([]int{2, 1}), false},
		{Ints2Tree([]int{1, 2, 1}), Ints2Tree([]int{1, 1, 2}), false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isSameTree(ts.p, ts.q))
		})
	}
}
