package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-cameras/

// Given a binary tree, we install cameras on the nodes of the tree.
// Each camera at a node can monitor its parent, itself, and its immediate children.
// Calculate the minimum number of cameras needed to monitor all nodes of the tree.

// Note:
// The number of nodes in the given tree will be in the range [1, 1000].
// Every node has value 0.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	result := 0
	covered := map[*TreeNode]bool{}
	covered[nil] = true

	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node, parent *TreeNode) {
		if node != nil {
			dfs(node.Left, node)
			dfs(node.Right, node)
			//
			if (parent == nil && !covered[node]) || !covered[node.Left] || !covered[node.Right] {
				result++
				covered[node] = true
				covered[parent] = true
				covered[node.Left] = true
				covered[node.Right] = true
			}
		}
	}

	dfs(root, nil)
	return result
}

func TestMinCameraCover(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{0, 0, null, 0, 0}), 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minCameraCover(ts.root))
		})
	}
}
