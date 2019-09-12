package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	temp := head.Next
	head.Next = swapPairs(head.Next.Next)
	temp.Next = head

	return temp
}

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		head   *ListNode
		output *ListNode
	}{
		{NewListNode([]int{1, 2, 3, 4}), NewListNode([]int{2, 1, 4, 3})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, swapPairs(ts.head))
		})
	}
}
