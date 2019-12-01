package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/boats-to-save-people/
func numRescueBoats(people []int, limit int) int {
	// 边界检查
	if people == nil || len(people) < 1 || len(people) > 50000 {
		return 0
	}
	// 排序数组
	sort.Ints(people)
	// 定义两个指针:
	// 1. people[i] + people[j] <= limit: i++, j--
	// 2. people[i] + people[j] > limit: j--
	i, j, res := 0, len(people)-1, 0
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
		}
		j--
		res++
	}
	// 返回结果
	return res
}

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		people []int
		limit  int
		output int
	}{
		{[]int{1, 2}, 3, 1},
		{[]int{3, 2, 2, 1}, 3, 3},
		{[]int{3, 5, 3, 4}, 5, 4},
		{
			[]int{2, 49, 10, 7, 11, 41, 47, 2, 22, 6, 13, 12, 33, 18, 10, 26, 2, 6, 50, 10},
			50,
			11,
		},
		{
			[]int{1, 3, 3, 4},
			5,
			3,
		},
		{
			[]int{2, 4},
			5,
			2,
		},
		{[]int{3, 2, 3, 2, 2}, 6, 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numRescueBoats(ts.people, ts.limit))
		})
	}
}
