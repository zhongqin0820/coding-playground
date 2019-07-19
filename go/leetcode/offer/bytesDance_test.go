package offer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 239. Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/
// 输出一个数组在指定窗口下的最大值数列-滑动窗口下的最大值数列
func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if k <= 1 {
		return nums
	}
	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	g := k - 1 // 比参考文章的分组少一个，可以减少 max 函数的调用，理论上可以加速。

	left := make([]int, size)
	for i := 0; i < size; i++ {
		if i%g == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
	}

	right := make([]int, size)
	// size-1 很可能不是那组的最后一个，需要单独列出
	right[size-1] = nums[size-1]
	for j := size - 2; j >= 0; j-- {
		if (j+1)%g == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(nums[j], right[j+1])
		}
	}

	res := make([]int, size-k+1)
	for i := 0; i <= size-k; i++ {
		// right[i] 中保存了 nums[i:g*(i/g+1)] 中的最大值
		// left[i+k-1] 中保存了 nums[g*(i/g+1):i+k] 中的最大值
		res[i] = max(right[i], left[i+k-1])
	}

	return res
}

// 滑动窗口问题
func TestBytesDance20190719(t *testing.T) {
	tests := []struct {
		arr    []int
		wsize  int
		output []int
	}{
		{[]int{13, 19, 7, 8, 19, 21, 4, 8, 33}, 3, []int{19, 19, 19, 21, 21, 21, 33}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1, -1}, 1, []int{1, -1}},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxSlidingWindow(ts.arr, ts.wsize))
		})
	}
}
