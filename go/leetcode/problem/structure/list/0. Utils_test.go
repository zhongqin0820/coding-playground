package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeNode(t *testing.T) {
	mTree := Ints2Tree([]int{1, 2, null, 3, 4})
	t.Run("PrintTreeNode", func(t *testing.T) {
		ast := assert.New(t)
		ast.Equal("1,2,null,3,4", mTree.PrintTreeNode())
	})
	t.Run("TreeNode2IntS", func(t *testing.T) {
		ast := assert.New(t)
		ast.Equal([]int{1, 2, null, 3, 4}, mTree.Tree2Ints())
	})
}
