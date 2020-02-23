package contest

import (
	"fmt"
	"math/bits"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/biweekly-contest-20/problems/sort-integers-by-the-number-of-1-bits/

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if a, b := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j])); a == b {
			return arr[i] < arr[j]
		} else {
			return a < b
		}
	})
	return arr
}

func TestSortByBits(t *testing.T) {
	tests := []struct {
		arr    []int
		output []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
		{[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}},
		{[]int{10000, 10000}, []int{10000, 10000}},
		{[]int{2, 3, 5, 7, 11, 13, 17, 19}, []int{2, 3, 5, 17, 7, 11, 13, 19}},
		{[]int{10, 100, 1000, 10000}, []int{10, 100, 10000, 1000}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sortByBits(ts.arr))
		})
	}
}
