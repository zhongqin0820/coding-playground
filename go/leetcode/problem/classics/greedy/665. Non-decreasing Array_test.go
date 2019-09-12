package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/non-decreasing-array/description/
func checkPossibility(nums []int) bool {
	for i, n := 1, len(nums); i < n; i++ {
		// 出现降序的情况
		if nums[i-1] > nums[i] {
			// 将nums[i-1] = nums[i]
			pre := helper665(nums)
			pre[i-1] = pre[i]
			// 将nums[i] = nums[i-1]
			next := helper665(nums)
			next[i] = next[i-1]
			// 判断这两种情况是否有存在严格升序的情况
			return sort.IsSorted(sort.IntSlice(pre)) ||
				sort.IsSorted(sort.IntSlice(next))
		}
	}
	return true
}

func helper665(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	return res
}

func TestCheckPossibility(t *testing.T) {
	tests := []struct {
		nums   []int
		output bool
	}{
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, checkPossibility(ts.nums))
		})
	}
}
