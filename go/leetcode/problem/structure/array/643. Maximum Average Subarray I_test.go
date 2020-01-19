package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) float64 {
	// 连续窗口下的最大平均数
	temp := 0
	for i := 0; i < k; i++ {
		temp += nums[i]
	}
	max, size := temp, len(nums)
	for i := k; i < size; i++ {
		temp = temp - nums[i-k] + nums[i]
		if max < temp {
			max = temp
		}
	}
	return float64(max) / float64(k)
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findMaxAverage(ts.nums, ts.k))
		})
	}
}
