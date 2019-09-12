package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

// Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	res := [][]int{}
	last, nlast := root, root
	level := 0

	var bfs func(*TreeNode)
	bfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		queue = append(queue, root)
		cur := []int{}
		for len(queue) != 0 {
			node := queue[0]
			queue = queue[1:]
			cur = append(cur, node.Val)
			//
			if node.Left != nil {
				queue = append(queue, node.Left)
				nlast = node.Left
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nlast = node.Right
			}
			if node == last {
				last = nlast
				level++
				res = append([][]int{cur}, res...)
				cur = []int{}
			}
		}
	}
	bfs(root)

	return res
}

func TestLevelOrderBottom(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output [][]int
	}{
		{Ints2Tree([]int{3, 9, 20, null, null, 15, 7}), [][]int{{15, 7}, {9, 20}, {3}}},
		{Ints2Tree([]int{}), [][]int{}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, levelOrderBottom(ts.root))
		})
	}
}
