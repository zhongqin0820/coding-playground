package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
	// 边缘检测
	if nums == nil || len(nums) == 0 {
		return false
	}
	// 题意：k距离内，存在两个数相等
	// |j-i| == k
	// 使用一个map保存nums[i]与i的下标，第二次出现的下标<=k则为true
	var m map[int]int = make(map[int]int, len(nums))
	for i, num := range nums {
		if m[num] != 0 && i-m[num]+1 <= k {
			return true
		}
		m[num] = i + 1
	}
	return false
}

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output bool
	}{
		{[]int{}, -1, false},
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, containsNearbyDuplicate(ts.nums, ts.k))
		})
	}
}
