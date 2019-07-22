package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-bottom-left-tree-value/

// Given a binary tree, find the leftmost value in the last row of the tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	last, nlast := root, root
	queue := []*TreeNode{root}
	cur := []int{}
	pre := []int{}
	// post order?
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
			pre = append([]int{}, cur...)
			cur = []int{}
		}
	}
	if len(pre) == 0 {
		return -1
	}
	return pre[0]
}

func findBottomLeftValueAdv(root *TreeNode) int {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]

		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	}
	return root.Val
}

func TestFindBottomLeftValue(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{2, 1, 3}), 1},
		{Ints2Tree([]int{1, 2, 3, 4, null, 5, 6, null, null, 7}), 7},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findBottomLeftValue(ts.root))
		})
		t.Run(fmt.Sprintf("Example %d Adv", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findBottomLeftValueAdv(ts.root))
		})
	}
}
