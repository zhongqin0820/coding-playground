package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/house-robber-iii/description/

// The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

// Determine the maximum amount of money the thief can rob tonight without alerting the police.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		la, lb := dfs(node.Left)
		ra, rb := dfs(node.Right)

		a := node.Val + lb + rb                    //rob node
		b := helper337(la, lb) + helper337(ra, rb) //not rob node, rob its sons
		return a, b
	}

	return helper337(dfs(root))
}

func helper337(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestRob(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{3, 2, 3, null, 3, null, 1}), 7},
		{Ints2Tree([]int{3, 4, 5, 1, 3, null, 1}), 9},
		{Ints2Tree([]int{4, 1, null, 2, null, 3}), 7},
		{Ints2Tree([]int{}), 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, rob(ts.root))
		})
	}
}
