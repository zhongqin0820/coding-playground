package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-169/problems/find-n-unique-integers-sum-up-to-zero
func sumZero(n int) []int {
	flag := n % 2
	num := -n / 2
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = num
		num++
		if flag == 0 && num == 0 {
			num++
		}
	}
	return res
}

func TestSumZero(t *testing.T) {
	tests := []struct {
		n      int
		output []int
	}{
		{5, []int{-2, -1, 0, 1, 2}},
		{3, []int{-1, 0, 1}},
		{1, []int{0}},
		{4, []int{-2, -1, 1, 2}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sumZero(ts.n))
		})
	}
}
