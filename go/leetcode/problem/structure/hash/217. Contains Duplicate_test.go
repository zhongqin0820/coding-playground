package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/contains-duplicate/description/
func containsDuplicate(nums []int) bool {
	var m map[int]struct{} = make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return ok
		}
		m[num] = struct{}{}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		output bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, containsDuplicate(ts.nums))
		})
	}
}
