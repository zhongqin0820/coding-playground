package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/third-maximum-number/
func thirdMax(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num == max1 || num == max2 { // 过滤重复的数
			continue
		}
		switch {
		case max1 < num:
			max3, max2, max1 = max2, max1, num
		case max2 < num:
			max3, max2 = max2, num
		case max3 < num:
			max3 = num
		}
	}
	if max3 == math.MinInt64 { // 如果没有第3大的数，返回第1大
		return max1
	}
	return max3
}

func TestThirdMax(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{2, 2, 2, 1}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, thirdMax(ts.nums))
		})
	}
}
