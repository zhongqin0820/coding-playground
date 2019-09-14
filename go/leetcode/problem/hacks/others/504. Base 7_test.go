package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/base-7/description/
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var sign string
	if num < 0 {
		sign = "-"
		num = -num
	}
	var s string
	for num > 0 {
		s = fmt.Sprintf("%d", num%7) + s
		num /= 7
	}
	return sign + s
}

func TestConvertToBase7(t *testing.T) {
	tests := []struct {
		num    int
		output string
	}{
		{100, "202"},
		{-7, "-10"},
		{0, "0"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, convertToBase7(ts.num))
		})
	}
}
