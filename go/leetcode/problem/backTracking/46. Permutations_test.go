package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/permutations/

// Given a collection of distinct integers, return all possible permutations.

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	cur := []int{}
	n := len(nums)
	visited := make([]bool, n)

	var bt func()
	bt = func() {
		if len(cur) == n {
			// 因为切片是引用的，因此需要使用深拷贝
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			cur = append(cur, nums[i])
			bt()
			cur = cur[:len(cur)-1]
			visited[i] = false
		}
	}
	bt()
	return res
}

func TestPermute(t *testing.T) {
	tests := []struct {
		nums   []int
		output [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, permute(ts.nums))
		})
	}
}
