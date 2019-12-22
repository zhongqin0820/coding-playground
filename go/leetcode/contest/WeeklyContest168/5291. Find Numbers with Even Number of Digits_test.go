package contest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-168/problems/find-numbers-with-even-number-of-digits/
func findNumbers(nums []int) int {
	res := 0
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			res++
		}
	}
	return res
}

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findNumbers(ts.nums))
		})
	}
}
