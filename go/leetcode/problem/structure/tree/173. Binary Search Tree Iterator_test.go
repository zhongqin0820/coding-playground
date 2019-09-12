package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-search-tree-iterator/

// Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
// Calling next() will return the next smallest number in the BST.

// Note:
// next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
// You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0, 128)
	res := BSTIterator{
		stack: stack,
	}
	res.push(root)
	return res
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	n := len(this.stack)
	res := this.stack[n-1]
	this.stack = this.stack[:n-1]
	this.push(res.Right)
	return res.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// push the left part of node
func (this *BSTIterator) push(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
func TestBSTIterator(t *testing.T) {
	ast := assert.New(t)
	var iterator BSTIterator = Constructor(Ints2Tree([]int{7, 3, 15, null, null, 9, 20}))
	// BSTIterator iterator = new BSTIterator(root);
	ast.Equal(iterator.Next(), 3)
	// iterator.next();    // return 3
	ast.Equal(iterator.Next(), 7)
	// iterator.next();    // return 7
	ast.Equal(iterator.HasNext(), true)
	// iterator.hasNext(); // return true
	ast.Equal(iterator.Next(), 9)
	// iterator.next();    // return 9
	ast.Equal(iterator.HasNext(), true)
	// iterator.hasNext(); // return true
	ast.Equal(iterator.Next(), 15)
	// iterator.next();    // return 15
	ast.Equal(iterator.HasNext(), true)
	// iterator.hasNext(); // return true
	ast.Equal(iterator.Next(), 20)
	// iterator.next();    // return 20
	ast.Equal(iterator.HasNext(), false)
	// iterator.hasNext(); // return false

}
