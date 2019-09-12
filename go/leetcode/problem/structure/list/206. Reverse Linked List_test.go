package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// head.Next指向prev负责反转
		// prev,head的操作负责移动指针
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		head   *ListNode
		output *ListNode
	}{
		{NewListNode([]int{1, 2, 3, 4, 5}), NewListNode([]int{5, 4, 3, 2, 1})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, reverseList(ts.head))
		})
	}
}
