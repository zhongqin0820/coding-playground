package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 获取 list 中的值
	nums := make([]int, 0, 64)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	// 按照规则对比值
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] != nums[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		head   *ListNode
		output bool
	}{
		{NewListNode([]int{1, 2}), false},
		{NewListNode([]int{1, 2, 2, 1}), true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isPalindrome(ts.head))
		})
	}
}

// func isPalindrome(head *ListNode) bool {
//     // base case
//     if head == nil || head.Next == nil {
//         return true
//     }

//     // 1. find the half
//     var prev *ListNode = nil
//     var slow *ListNode = head
//     var fast *ListNode = head

//     for fast != nil && fast.Next != nil {
//         prev = slow
//         slow = slow.Next
//         fast = fast.Next.Next
//     }
//     // cut half
//     prev.Next = nil

//     // 2. set the second half and reverse the first half
//     if fast != nil {
//         slow = slow.Next
//     }
//     var head2 *ListNode = slow
//     head = reverse(head)

//     // 3. compare two linked lists
//     for head != nil && head2 != nil {
//         if head.Val != head2.Val {
//             return false
//         }
//         head = head.Next
//         head2 = head2.Next
//     }
//     return true;
// }

// func reverse(head *ListNode) *ListNode {
//     var prev *ListNode = nil
//     var curr *ListNode = head

//     for curr != nil {
//         var post *ListNode = curr.Next
//         curr.Next = prev
//         prev = curr
//         curr = post
//     }

//     return prev
// }
