package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/permutations-ii/description/
func permuteUnique(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	n := len(nums)
	vector := make([]int, n)
	visited := make([]bool, n)
	var res [][]int
	helper47(0, n, nums, vector, visited, &res)
	return res
}

func helper47(cur, n int, nums, vector []int, visited []bool, ans *[][]int) {
	if cur == n {
		temp := make([]int, n)
		copy(temp, vector)
		*ans = append(*ans, temp)
		return
	}
	used := make(map[int]bool, n-cur)
	for i := 0; i < n; i++ {
		if !visited[i] && !used[nums[i]] {
			used[nums[i]] = true
			visited[i] = true
			vector[cur] = nums[i]
			helper47(cur+1, n, nums, vector, visited, ans)
			visited[i] = false
		}
	}
}

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		nums   []int
		output [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, permuteUnique(ts.nums))
		})
	}
}
