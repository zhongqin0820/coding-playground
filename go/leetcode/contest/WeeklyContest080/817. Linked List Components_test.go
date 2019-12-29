package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-80/problems/linked-list-components/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	// 使用一个map存储G中的节点值
	isInG := make(map[int]bool, len(G))
	for _, g := range G {
		isInG[g] = true
	}
	res := 0
	for ; head != nil; head = head.Next {
		// 遍历到节点的下一个节点为空 || 遍历到节点的下一个节点的值不在G中，则res ++
		if isInG[head.Val] &&
			(head.Next == nil || !isInG[head.Next.Val]) {
			res++
		}
	}
	return res
}

func TestNumComponents(t *testing.T) {
	tests := []struct {
		head   *ListNode
		G      []int
		output int
	}{
		{NewListNode([]int{0, 1, 2, 3}), []int{0, 1, 3}, 2},
		{NewListNode([]int{0, 1, 2, 3, 4}), []int{0, 3, 1, 4}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numComponents(ts.head, ts.G))
		})
	}
}
