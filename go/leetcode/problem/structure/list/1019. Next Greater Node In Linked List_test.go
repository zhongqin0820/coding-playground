package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/next-greater-node-in-linked-list/

// We are given a linked list with head as the first node.  Let's number the nodes in the list: node_1, node_2, node_3, ... etc.

// Each node may have a next larger value: for node_i, next_larger(node_i) is the node_j.val such that j > i, node_j.val > node_i.val, and j is the smallest possible choice.  If such a j does not exist, the next larger value is 0.

// Return an array of integers answer, where answer[i] = next_larger(node_{i+1}).

// Note that in the example inputs (not outputs) below, arrays such as [2,1,5] represent the serialization of a linked list with a head node value of 2, second node value of 1, and third node value of 5.
// - 1 <= node.val <= 10^9 for each node in the linked list.
// - The given list has length in the range [0, 10000].

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	A := helper1019(head)
	res := make([]int, len(A))
	stack, peek := make([]int, len(A)), -1
	for i := 0; i < len(A); i++ {
		for peek >= 0 && A[stack[peek]] < A[i] {
			res[stack[peek]] = A[i]
			peek--
		}
		peek++
		stack[peek] = i
	}
	return res
}

func helper1019(head *ListNode) []int {
	res := make([]int, 0, 1024)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestNextLargerNodes(t *testing.T) {
	tests := []struct {
		input  *ListNode
		output []int
	}{
		{NewListNode([]int{2, 1, 5}), []int{5, 5, 0}},
		{NewListNode([]int{2, 7, 4, 3, 5}), []int{7, 0, 5, 5, 0}},
		{NewListNode([]int{1, 7, 5, 1, 9, 2, 5, 1}), []int{7, 9, 9, 9, 0, 5, 0, 0}},
		{NewListNode([]int{3, 3}), []int{0, 0}},
		{NewListNode([]int{4, 3, 2, 5, 1, 8, 10}), []int{5, 5, 5, 8, 8, 10, 0}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, nextLargerNodes(ts.input))
		})
	}
}
