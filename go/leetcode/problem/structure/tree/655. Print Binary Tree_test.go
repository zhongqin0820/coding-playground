package problem

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/print-binary-tree/

// Print a binary tree in an m*n 2D string array following these rules:
// The row number m should be equal to the height of the given binary tree.
// The column number n should always be an odd number.
// The root node's value (in string format) should be put in the exactly middle of the first row it can be put. The column and the row where the root node belongs will separate the rest space into two parts (left-bottom part and right-bottom part). You should print the left subtree in the left-bottom part and print the right subtree in the right-bottom part. The left-bottom part and the right-bottom part should have the same size. Even if one subtree is none while the other is not, you don't need to print anything for the none subtree but still need to leave the space as large as that for the other subtree. However, if two subtrees are none, then you don't need to leave space for both of them.
// Each unused space should contain an empty string "".
// Print the subtrees following the same rules.

// Note: The height of binary tree is in the range of [1, 10].

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	m := helper655GetLevel(root) // m=height of the given tree
	n := 1<<uint(m) - 1          //n=2^m-1
	res := make([][]string, m)
	for i := 0; i < m; i++ {
		res[i] = make([]string, n)
	}
	loc := 1<<uint(m-1) - 1
	helper655(root, 0, loc, &res)
	return res
}

func helper655(root *TreeNode, i, j int, res *[][]string) {
	if root == nil {
		return
	}
	(*res)[i][j] = strconv.Itoa(root.Val)
	m := len(*res)
	if m-i-2 < 0 {
		return
	}
	diff := 1 << uint(m-i-2)
	helper655(root.Left, i+1, j-diff, res)
	helper655(root.Right, i+1, j+diff, res)
}

func helper655GetLevel(root *TreeNode) (res int) {
	res = 1
	queue := []*TreeNode{root}
	for {
		hasChild := false
		// 遍历到有孩子的节点，如果没有找到break返回层数
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil || queue[i].Right != nil {
				hasChild = true
				break
			}
		}

		if !hasChild {
			break
		}
		res++
		// 迭代有孩子的部分节点
		fQueue := queue[:len(queue)]
		queue = queue[len(queue):]

		for i := 0; i < len(fQueue); i++ {
			if fQueue[i].Left != nil {
				queue = append(queue, fQueue[i].Left)
			}
			if fQueue[i].Right != nil {
				queue = append(queue, fQueue[i].Right)
			}
		}
	}
	return res
}

func TestPrintTree(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output [][]string
	}{
		{Ints2Tree([]int{1, 2}), [][]string{[]string{"", "1", ""}, []string{"2", "", ""}}},
		{Ints2Tree([]int{1, 2, 3, null, 4}), [][]string{[]string{"", "", "", "1", "", "", ""}, []string{"", "2", "", "", "", "3", ""}, []string{"", "", "4", "", "", "", ""}}},
		{Ints2Tree([]int{1, 2, 5, 3, null, null, null, 4}), [][]string{[]string{"", "", "", "", "", "", "", "1", "", "", "", "", "", "", ""}, []string{"", "", "", "2", "", "", "", "", "", "", "", "5", "", "", ""}, []string{"", "3", "", "", "", "", "", "", "", "", "", "", "", "", ""}, []string{"4", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, printTree(ts.root))
		})
	}
}
