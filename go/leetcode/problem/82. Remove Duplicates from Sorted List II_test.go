package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

// Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

// Time: O(n)
// Space: O(n)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 结果的初始节点
	res := &ListNode{
		Next: nil,
		Val:  math.MaxInt64,
	}
	ret, cur := res, head
	preVal := ret.Val
	for cur != nil {
		// 如果当前节点不与前面节点重复
		if cur.Val != preVal {
			// 且当前节点不与后面节点重复，则加入结果节点list中
			if cur.Next == nil || cur.Next.Val != cur.Val {
				res.Next = &ListNode{
					Next: nil,
					Val:  cur.Val,
				}
				res = res.Next
			}
		}
		// 遍历
		preVal = cur.Val
		cur = cur.Next
	}
	return ret.Next
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		input, output *ListNode
	}{
		{nil, nil},
		{NewListNode([]int{1, 2, 3, 3, 4, 4, 5}), NewListNode([]int{1, 2, 5})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, deleteDuplicates(ts.input))
			ast.Equal(ts.output.PrintList(), deleteDuplicates(ts.input).PrintList())
		})
	}
}
