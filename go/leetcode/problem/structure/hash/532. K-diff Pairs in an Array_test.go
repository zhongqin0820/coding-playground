package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/k-diff-pairs-in-an-array
func findPairs(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k < 0 {
		return 0
	}
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	res := 0
	if k == 0 {
		for _, cnt := range freq {
			if cnt > 1 {
				res++
			}
		}
	} else {
		for num, _ := range freq {
			if freq[num-k] > 0 {
				res++
			}
		}
	}
	return res
}

func TestFindPairs(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output int
	}{
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
		{[]int{1, 3, 1, 5, 4}, 0, 1},
		{[]int{3, 3, 5, 5, 5}, 2, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findPairs(ts.nums, ts.k))
		})
	}
}
