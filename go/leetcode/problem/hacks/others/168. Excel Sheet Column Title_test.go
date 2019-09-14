package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/excel-sheet-column-title/description/
func convertToTitle(n int) string {
	var m []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var s string
	for n > 0 {
		// 因为是从 1 开始计算的，而不是从 0 开始，因此需要对 n 执行 -1 操作
		n--
		s = m[n%26] + s
		n /= 26
	}
	return s
}

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		n      int
		output string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, convertToTitle(ts.n))
		})
	}
}
