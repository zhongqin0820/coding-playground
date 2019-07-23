package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

// Given a non-empty special binary tree consisting of nodes with the non-negative value, where each node in this tree has exactly two or zero sub-node. If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes. More formally, the property root.val = min(root.left.val, root.right.val) always holds.

// Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.

// If no such second minimum value exists, output -1 instead.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	res := math.MaxInt64
	val := root.Val

	var pre func(*TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		if val < node.Val && node.Val < res {
			res = node.Val
		}
		pre(node.Left)
		pre(node.Right)
	}

	pre(root)
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

func TestFindSecondMinimumValue(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{2, 2, 5, null, null, 5, 7}), 5},
		{Ints2Tree([]int{2, 2, 2}), -1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findSecondMinimumValue(ts.root))
		})
	}
}
