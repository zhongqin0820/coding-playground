package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/

// Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return helper109(head, nil)
}
func helper109(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	slow, fast := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return &TreeNode{
		Val:   slow.Val,
		Left:  helper109(begin, slow),
		Right: helper109(slow.Next, end),
	}
}

func TestSortedListToBST(t *testing.T) {
	tests := []struct {
		head   *ListNode
		output *TreeNode
	}{
		{NewListNode([]int{-10, -3, 0, 5, 9}), Ints2Tree([]int{0, -3, 9, -10, null, 5})},
		{nil, nil},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sortedListToBST(ts.head))
		})
	}
}
