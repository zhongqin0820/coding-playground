package problem

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// new a link list via a 1D slice
func NewListNode(a []int) *ListNode {
	if a == nil || len(a) == 0 {
		return nil
	}
	//
	node := new(ListNode)
	node.Val = a[0]
	head := node
	for i := 1; i < len(a); i++ {
		temp := &ListNode{
			Val:  a[i],
			Next: nil,
		}
		node.Next = temp
		node = node.Next
	}
	return head
}

// print the link list for test usage
func (this *ListNode) PrintList() string {
	that := this
	res := ""
	for that != nil && that.Next != nil {
		res += fmt.Sprintf("%d->", that.Val)
		that = that.Next
	}
	if that != nil {
		res += fmt.Sprintf("%d", that.Val)
	}
	return res
}

// compare two linked list
func (this *ListNode) Compare(node *ListNode) bool {
	that := this
	for node != nil && that != nil {
		if node.Val != that.Val {
			return false
		}
		node = node.Next
		that = that.Next
	}
	return node == nil && that == nil
}

// null indicates null node in the Tree
var null int = -1 << 63

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Ints2Tree(a []int) *TreeNode {
	if a == nil || len(a) == 0 {
		return nil
	}
	n := len(a)
	root := &TreeNode{Val: a[0]}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	for i := 1; i < n; {
		node := queue[0]
		queue = queue[1:]
		//
		if i < n && a[i] != null {
			node.Left = &TreeNode{Val: a[i]}
			queue = append(queue, node.Left)
		}
		i++
		//
		if i < n && a[i] != null {
			node.Right = &TreeNode{Val: a[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func PreIn2Tree(pre, in []int) *TreeNode {
	if pre == nil || in == nil || len(pre) != len(in) || len(pre) == 0 {
		return nil
	}
	res := &TreeNode{Val: pre[0]}
	if len(in) == 1 {
		return res
	}

	idx := -1
	for i, v := range in {
		if idx == v {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil
	}

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func (tn *TreeNode) Tree2Ints() []int {
	res := make([]int, 0, 1024)
	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			if node == nil {
				res = append(res, null)
			} else {
				res = append(res, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[n:]
	}

	i := len(res)
	for i > 0 && res[i-1] == null {
		i--
	}
	return res[:i]
}

func (tn *TreeNode) PrintTreeNode() string {
	res := tn.Tree2Ints()
	ret := []byte{}
	for j := 0; j < len(res); j++ {
		if res[j] == null {
			ret = append(ret, []byte("null")...)
		} else {
			ret = strconv.AppendInt(ret, int64(res[j]), 10)
		}
		ret = append(ret, ',')
	}
	return strings.TrimRight(string(ret), ",")
}

// compare two trees
func (tn *TreeNode) Compare(a *TreeNode) bool {
	if tn == nil && a == nil {
		return true
	}
	if tn == nil || a == nil || tn.Val != a.Val {
		return false
	}
	return tn.Left.Compare(a.Left) && tn.Right.Compare(a.Right)
}

// compare two 2D slice
func slice2DCompare(a, b *[][]int) bool {
	if a == nil || b == nil || *a == nil || *b == nil {
		return false
	}
	if len((*a)) == 0 && len((*a)) == len((*b)) {
		return true
	}
	if len((*a)) != len((*b)) || len((*a)[0]) != len((*b)[0]) {
		return false
	}
	for i, v := range *a {
		for j, val := range (*b)[i] {
			if v[j] != val {
				return false
			}
		}
	}
	return true
}

// compare two 1D slice
func slice1DCompare(a, b []int) bool {
	if a == nil || b == nil || len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(a) == len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// compare two 1D string
func string1DCompare(a, b []string) bool {
	if a == nil || b == nil || len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(a) == len(b) {
		return true
	}
	for i, v := range a {
		if strings.Compare(v, b[i]) != 0 {
			return false
		}
	}
	return true
}
