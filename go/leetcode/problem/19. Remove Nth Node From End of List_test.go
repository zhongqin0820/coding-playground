package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// Given a linked list, remove the n-th node from the end of list and return its head.

// Given linked list: 1->2->3->4->5, and n = 2.
// After removing the second node from the end, the linked list becomes 1->2->3->5.

// Note:
// Given n will always be valid.

// Follow up:
// Could you do this in one pass?

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := new(ListNode)
	ret, cur, curK := res, head, head
	ret.Next = head
	//
	for i := 1; i < n && curK != nil; i++ {
		// log.Println(curK.Val)
		curK = curK.Next
	}
	//
	for curK.Next != nil {
		ret = ret.Next
		cur = cur.Next
		curK = curK.Next
	}
	//
	ret.Next = cur.Next
	return res.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		head   *ListNode
		n      int
		output *ListNode
	}{
		{NewListNode([]int{1, 2, 3, 4, 5}), 2, NewListNode([]int{1, 2, 3, 5})},
		{NewListNode([]int{1, 2, 3, 4, 5}), 1, NewListNode([]int{1, 2, 3, 4})},
		{NewListNode([]int{1, 2, 3, 4, 5}), 5, NewListNode([]int{2, 3, 4, 5})},
		{NewListNode([]int{1}), 1, NewListNode([]int{})},
		{NewListNode([]int{1, 2}), 1, NewListNode([]int{1})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintList(), removeNthFromEnd(ts.head, ts.n).PrintList())
		})
	}
}
