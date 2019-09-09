package problem

// https://leetcode.com/problems/unique-binary-search-trees-ii/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return helper95(1, n)
}

func helper95(start, end int) []*TreeNode {
	list := make([]*TreeNode, 0)
	if start > end {
		list = append(list, nil)
		return list
	}
	var left, right []*TreeNode
	for i := start; i <= end; i++ {
		left = helper95(start, i-1)
		right = helper95(i+1, end)
		for _, leftV := range left {
			for _, rightV := range right {
				root := &TreeNode{Val: i}
				root.Left = leftV
				root.Right = rightV
				list = append(list, root)
			}
		}
	}
	return list
}
