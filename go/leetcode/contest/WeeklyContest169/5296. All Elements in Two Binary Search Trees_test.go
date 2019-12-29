package contest

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type intHeap []int

func (hs intHeap) Len() int { return len(hs) }

func (hs intHeap) Swap(i, j int) { hs[i], hs[j] = hs[j], hs[i] }

func (hs intHeap) Less(i, j int) bool { return hs[i] < hs[j] }

func (hs *intHeap) Push(e interface{}) { *hs = append(*hs, e.(int)) }

func (hs *intHeap) Pop() interface{} {
	e := (*hs)[hs.Len()-1]
	*hs = (*hs)[:hs.Len()-1]
	return e
}

// https://leetcode.com/contest/weekly-contest-169/problems/all-elements-in-two-binary-search-trees
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	h := &intHeap{}
	heap.Init(h)
	helper5296(root1, h)
	helper5296(root2, h)
	res := make([]int, 0, h.Len())
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

func helper5296(root *TreeNode, h *intHeap) {
	if root == nil {
		return
	}
	helper5296(root.Left, h)
	heap.Push(h, root.Val)
	helper5296(root.Right, h)
}

func TestGetAllElements(t *testing.T) {
	tests := []struct {
		root1  *TreeNode
		root2  *TreeNode
		output []int
	}{
		{Ints2Tree([]int{2, 1, 4}), Ints2Tree([]int{1, 0, 3}), []int{0, 1, 1, 2, 3, 4}},
		{Ints2Tree([]int{0, -10, 10}), Ints2Tree([]int{5, 1, 7, 0, 2}), []int{-10, 0, 0, 1, 2, 5, 7, 10}},
		{Ints2Tree([]int{}), Ints2Tree([]int{5, 1, 7, 0, 2}), []int{0, 1, 2, 5, 7}},
		{Ints2Tree([]int{0, -10, 10}), Ints2Tree([]int{}), []int{-10, 0, 10}},
		{Ints2Tree([]int{1, null, 8}), Ints2Tree([]int{8, 1}), []int{1, 1, 8, 8}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, getAllElements(ts.root1, ts.root2))
		})
	}
}
