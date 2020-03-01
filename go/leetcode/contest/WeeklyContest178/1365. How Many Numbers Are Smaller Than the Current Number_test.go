package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	temp := make([]int, n)
	copy(temp, nums)
	sort.Ints(temp)
	m := make(map[int]int, n)
	for i := n - 1; i >= 0; i-- {
		m[temp[i]] = i
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = m[nums[i]]
	}
	return res
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
		{[]int{7, 7, 7, 7}, []int{0, 0, 0, 0}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, smallerNumbersThanCurrent(ts.nums))
		})
	}
}
