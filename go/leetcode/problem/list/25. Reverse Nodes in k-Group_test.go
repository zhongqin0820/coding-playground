package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-nodes-in-k-group/

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

// Example:
// Given this linked list: 1->2->3->4->5
// For k = 2, you should return: 2->1->4->3->5
// For k = 3, you should return: 3->2->1->4->5

// Note:
// Only constant extra memory is allowed.
// You may not alter the values in the list's nodes, only nodes itself may be changed.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}
	// 获得[i,i+k]的尾节点
	tail, ok := helper25GetTail(head, k)
	// ok==true需要翻转
	if ok {
		// 保存tail.Next
		tailNext := tail.Next
		tail.Next = nil
		// 翻转[i,i+k]
		head, tail = helper25ReverseK(head, tail)
		// 递归翻转剩下的部分
		tail.Next = reverseKGroup(tailNext, k)
	}
	return head
}

func helper25GetTail(head *ListNode, k int) (*ListNode, bool) {
	for head != nil && k > 1 {
		k--
		head = head.Next
	}
	return head, k == 1 && head != nil
}

func helper25ReverseK(head, tail *ListNode) (*ListNode, *ListNode) {
	curPre, cur := head, head.Next
	for cur != nil {
		curPre, cur, cur.Next = cur, cur.Next, curPre
	}
	return tail, head
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		head   *ListNode
		k      int
		output *ListNode
	}{
		{NewListNode([]int{1, 2, 3, 4, 5}), 2, NewListNode([]int{2, 1, 4, 3, 5})},
		{NewListNode([]int{1, 2, 3, 4, 5}), 3, NewListNode([]int{3, 2, 1, 4, 5})},
		{NewListNode([]int{1, 2, 3, 4, 5}), 1, NewListNode([]int{1, 2, 3, 4, 5})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintList(), reverseKGroup(ts.head, ts.k).PrintList())
		})
	}
}
