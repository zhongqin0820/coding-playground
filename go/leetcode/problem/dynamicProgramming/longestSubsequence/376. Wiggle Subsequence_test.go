package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/wiggle-subsequence/description/

// A sequence of numbers is called a wiggle sequence if **the differences between successive numbers strictly alternate between positive and negative.** The first difference (if one exists) may be either positive or negative. A sequence with fewer than two elements is trivially a wiggle sequence.

// For example, [1,7,4,9,2,5] is a wiggle sequence because the differences (6,-3,5,-7,3) are alternately positive and negative. In contrast, [1,4,7,2,5] and [1,7,4,5,5] are not wiggle sequences, the first because its first two differences are positive and the second because its last difference is zero.

// Given a sequence of integers, return the length of the longest subsequence that is a wiggle sequence. A subsequence is obtained by deleting some number of elements (eventually, also zero) from the original sequence, leaving the remaining elements in their original order.

// Follow up:
// Can you do it in O(n) time?

func wiggleMaxLength(nums []int) int {
	// edge case
	switch {
	case nums == nil:
		return 0
	case len(nums) < 2:
		return len(nums)
	}
	res, n := 0, len(nums)
	//
	checkFunc := func(n int) func(int) {
		init := 1
		if n < 0 {
			init = -1
		}
		res = 2
		return func(x int) {
			var cur int
			switch {
			case x < 0:
				cur = -1
			case x > 0:
				cur = 1
			default:
				return
			}
			if init*cur < 0 {
				res++
				init = cur
			}
		}
	}
	var check func(int)

	var i = 1
	for i < n && nums[i]-nums[i-1] == 0 {
		i++
	}
	if i == n {
		return 1
	}
	check = checkFunc(nums[i] - nums[i-1])
	for ; i < n; i++ {
		check(nums[i] - nums[i-1])
	}
	return res
}

// ref: https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92.md#3-%E6%9C%80%E9%95%BF%E6%91%86%E5%8A%A8%E5%AD%90%E5%BA%8F%E5%88%97
func wiggleMaxLengthAdv(nums []int) int {
	// edge case
	switch {
	case nums == nil:
		return 0
	case len(nums) < 2:
		return len(nums)
	}
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] { //正号
			up = down + 1
		} else if nums[i] < nums[i-1] { //负号
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}

func TestWiggleMaxLength(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{0, 0}, 1},
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		// Explanation: The entire sequence is a wiggle sequence.
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		// Explanation: There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, wiggleMaxLength(ts.nums))
			ast.Equal(ts.output, wiggleMaxLengthAdv(ts.nums))
		})
	}
}
