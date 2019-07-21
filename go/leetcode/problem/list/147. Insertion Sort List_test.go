package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/insertion-sort-list/

// Sort a linked list using insertion sort.

// Example 1:

// Input: 4->2->1->3
// Output: 1->2->3->4
// Example 2:

// Input: -1->5->3->4->0
// Output: -1->0->3->4->5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res, cur := &ListNode{Next: head}, head
	//
	for cur != nil && cur.Next != nil {
		node := cur.Next
		if node.Val > cur.Val {
			cur = node
			continue
		}
		// delete node
		cur.Next = node.Next
		// insert node by doing insertion
		pre, preNext := res, res.Next
		for preNext.Val < node.Val {
			pre = preNext
			preNext = preNext.Next
		}
		pre.Next = node
		node.Next = preNext
	}
	return res.Next
}

func TestInsertionSortList(t *testing.T) {
	tests := []struct {
		input, output *ListNode
	}{
		{NewListNode([]int{4, 2, 1, 3}), NewListNode([]int{1, 2, 3, 4})},
		{NewListNode([]int{-1, 5, 3, 4, 0}), NewListNode([]int{-1, 0, 3, 4, 5})},
		{nil, nil},
		{NewListNode([]int{}), NewListNode([]int{})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintList(), insertionSortList(ts.input).PrintList())
		})
	}
}
