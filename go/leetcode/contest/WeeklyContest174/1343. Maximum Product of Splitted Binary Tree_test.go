package contest

import (
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
var mod = 1000000007

func maxProduct(root *TreeNode) int {
	// 拷贝并计算当前子树和
	temp, res := sum(root)
	// max_child (root.Val - child.Val)*(child.Val)
	res = pos(temp, temp)
	return res % mod
}

func pos(node, root *TreeNode) int {
	if node == nil {
		return 0
	}
	l := pos(node.Left, root)
	r := pos(node.Right, root)
	return max3(l, max3(r, (root.Val-node.Val)*node.Val))
}

func max3(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func sum(node *TreeNode) (*TreeNode, int) {
	if node == nil {
		return node, 0
	}
	temp, l, r := new(TreeNode), 0, 0
	temp.Left, l = sum(node.Left)
	temp.Right, r = sum(node.Right)
	temp.Val = node.Val + l + r
	return temp, temp.Val
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{1, 2, 3, 4, 5, 6}), 110},
		{Ints2Tree([]int{1, null, 2, 3, 4, null, null, 5, 6}), 90},
		{Ints2Tree([]int{2, 3, 9, 10, 7, 8, 6, 5, 4, 11, 1}), 1025},
		{Ints2Tree([]int{1, 1}), 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxProduct(ts.root))
		})
	}
}
