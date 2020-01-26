package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/jump-game-ii/
// ref: https://www.cnblogs.com/grandyang/p/4373533.html
func jump(nums []int) int {
	res, size, last, cur := 0, len(nums), 0, 0
	for i := 0; i < size-1; i++ {
		// 判断是否最后可达
		if cur < i+nums[i] {
			cur = i + nums[i]
		}
		// 更新跳数
		if i == last {
			last = cur
			res++
			if cur >= size-1 {
				break
			}
		}
	}
	return res
}

func TestJump(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, jump(ts.nums))
		})
	}
}
