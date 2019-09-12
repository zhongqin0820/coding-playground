package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

// Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2

	return &TreeNode{
		Left:  sortedArrayToBST(nums[:m]),
		Val:   nums[m],
		Right: sortedArrayToBST(nums[m+1:]),
	}
}

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums   []int
		output *TreeNode
	}{
		{[]int{-10, -3, 0, 5, 9}, Ints2Tree([]int{0, -3, 9, -10, null, 5})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sortedArrayToBST(ts.nums))
		})
	}
}
