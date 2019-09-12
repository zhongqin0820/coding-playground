package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/path-sum-ii/

// Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

// Note: A leaf is a node with no children.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		//
		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		sum -= node.Val

		//
		if node.Left == nil && node.Right == nil && sum == 0 {
			temp := make([]int, level+1)
			copy(temp, path)
			res = append(res, temp)
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}
	dfs(root, 0, sum)
	return res
}

func TestPathSum(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		sum    int
		output [][]int
	}{
		{Ints2Tree([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1}), 22, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, pathSum(ts.root, ts.sum))
		})
	}
}
