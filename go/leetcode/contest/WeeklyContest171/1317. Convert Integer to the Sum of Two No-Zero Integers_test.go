package contest

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-171/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
func getNoZeroIntegers(n int) []int {
	root := n / 2
	for A := 1; A <= root; A++ {
		if strings.Contains(strconv.Itoa(A), "0") {
			continue
		}
		B := n - A
		if strings.Contains(strconv.Itoa(B), "0") {
			continue
		}
		return []int{A, B}
	}
	return []int{}
}

func TestGetNoZeroIntegers(t *testing.T) {
	tests := []struct {
		n      int
		output []int
	}{
		// 2 <= n <= 10^4
		{2, []int{1, 1}},
		{11, []int{2, 9}},
		{10000, []int{1, 9999}},
		{69, []int{1, 68}},
		{1010, []int{11, 999}},
		{4, []int{1, 3}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, getNoZeroIntegers(ts.n))
		})
	}
}
