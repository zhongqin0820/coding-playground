package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-two-numbers/

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	curVal, carry := 0, 0
	res := new(ListNode)
	node := res
	//正常情况
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		cur := new(ListNode)
		curVal = l1.Val + l2.Val + carry
		cur.Val, carry = curVal%10, curVal/10
		node.Next = cur
		node = node.Next
	}
	//处理不等长情况
	if l2 != nil {
		l1 = l2
	}
	for ; l1 != nil; l1 = l1.Next {
		cur := new(ListNode)
		curVal = l1.Val + carry
		cur.Val, carry = curVal%10, curVal/10
		node.Next = cur
		node = node.Next
	}
	//处理carry不为0
	if carry > 0 {
		cur := new(ListNode)
		cur.Val = carry
		node.Next = cur
		node = node.Next
	}
	return res.Next
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1, l2 *ListNode
		output *ListNode
	}{
		{nil, nil, nil},
		{NewListNode([]int{2, 4, 3}), nil, NewListNode([]int{2, 4, 3})},
		{nil, NewListNode([]int{5, 6, 4}), NewListNode([]int{5, 6, 4})},
		{NewListNode([]int{2, 4, 3}), NewListNode([]int{5, 6, 4}), NewListNode([]int{7, 0, 8})},
		// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
		// Output: 7 -> 0 -> 8
		// Explanation: 342 + 465 = 807.
		{NewListNode([]int{2, 4, 9}), NewListNode([]int{5, 6, 4}), NewListNode([]int{7, 0, 4, 1})},
		{NewListNode([]int{9, 8, 7, 6, 5}), NewListNode([]int{1, 1, 2, 3, 4}), NewListNode([]int{0, 0, 0, 0, 0, 1})},
		{NewListNode([]int{1}), NewListNode([]int{9, 9}), NewListNode([]int{0, 0, 1})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintList(), addTwoNumbers(ts.l1, ts.l2).PrintList())
		})
	}
}
