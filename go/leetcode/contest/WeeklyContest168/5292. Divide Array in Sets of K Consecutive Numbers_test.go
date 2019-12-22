package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-168/problems/divide-array-in-sets-of-k-consecutive-numbers/
func isPossibleDivide(nums []int, k int) bool {
	if nums == nil || len(nums) == 0 || len(nums)%k != 0 {
		return false
	}
	// 记录<num:freq>以及初始idx
	m := make(map[int]int, len(nums))
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
			if minNum > num {
				minNum = num
			}
			if maxNum < num {
				maxNum = num
			}
		} else {
			m[num]++
		}
	}
	for len(m) != 0 {
		// 序列
		for k_ := 0; k_ < k; k_++ {
			num := minNum + k_
			if count, ok := m[num]; !ok || count <= 0 {
				return false
			}
			m[num]--
			if m[num] == 0 {
				delete(m, num)
			}
		}
		// 重新定位起始idx
		minNum = maxNum
		for num, count := range m {
			if num < minNum && count != 0 {
				minNum = num
			}
		}
	}
	return true
}

func TestIsPossibleDivide(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output bool
	}{
		{[]int{1, 2, 3, 3, 4, 4, 5, 6}, 4, true},
		{[]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3, true},
		{[]int{3, 3, 2, 2, 1, 1}, 3, true},
		{[]int{1, 2, 3, 4}, 3, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isPossibleDivide(ts.nums, ts.k))
		})
	}
}
