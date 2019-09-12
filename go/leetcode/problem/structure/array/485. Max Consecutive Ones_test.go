package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-consecutive-ones/description/

// Given a binary array, find the maximum number of consecutive 1s in this array.
// Note:
// The input array will only contain 0 and 1.
// The length of input array is a positive integer and will not exceed 10,000

func findMaxConsecutiveOnes(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	res := 0
	// 标记
	for i, j := 0, -1; i < len(nums); i++ {
		if nums[i] == 0 { //记录
			j = i
		} else {
			if res < i-j { //比较
				res = i - j
			}
		}
	}
	return res
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1, 0, 1, 1, 1}, 3},
		// Explanation: The first two digits or the last three digits are consecutive 1s.The maximum number of consecutive 1s is 3.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findMaxConsecutiveOnes(ts.nums))
		})
	}
}
