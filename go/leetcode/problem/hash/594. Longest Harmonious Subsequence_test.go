package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-harmonious-subsequence/description/
func findLHS(nums []int) int {
	// m统计对应值的数量
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}
	// 统计连续序列的最大长度
	// 注意序列不要求是数组中连续
	res := 0
	for num, count1 := range m {
		// 下一个连续的值是否存在
		if count2, ok := m[num+1]; ok {
			temp := count1 + count2
			if res < temp {
				res = temp
			}
		}
	}
	return res
}

func TestFindLHS(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findLHS(ts.nums))
		})
	}
}
