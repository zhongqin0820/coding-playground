package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-tilt/

// Given a binary tree, return the tilt of the whole tree.

// The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.
// The tilt of the whole tree is defined as the sum of all nodes' tilt.

// Note:
// The sum of node values in any subtree won't exceed the range of 32-bit integer.
// All the tilt values won't exceed the range of 32-bit integer.
//
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	var tilt int
	var pos func(*TreeNode) int
	pos = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := pos(node.Left)
		r := pos(node.Right)
		if l < r {
			tilt += r - l
		} else {
			tilt += l - r

		}
		return l + r + node.Val
	}
	pos(root)
	return tilt
}

func TestFindTilt(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{1, 2, 3}), 1},
		{Ints2Tree([]int{1}), 0},
		{Ints2Tree([]int{1, 2, 2}), 0},
		{Ints2Tree([]int{1, 2, 2, 3}), 6},
		{Ints2Tree([]int{1, 2, 3, 4, null, 5}), 11},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findTilt(ts.root))
		})
	}
}
