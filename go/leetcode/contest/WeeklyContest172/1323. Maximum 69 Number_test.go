package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-172/problems/maximum-69-number/
func maximum69Number(num int) int {
	exp := 0
	for i := 1; i < num; i *= 10 {
		if (num/i)%10 == 6 {
			exp = i
		}
	}
	if exp == 0 {
		return num
	}
	return (num/exp)*exp + 3*exp + num%exp

}

func TestMaximum69Number(t *testing.T) {
	tests := []struct {
		num    int
		output int
	}{
		{9669, 9969},
		{9996, 9999},
		{9999, 9999},
		{6, 9},
		{9, 9},
		{69, 99},
		{96, 99},
		{99, 99},
		{669, 969},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maximum69Number(ts.num))
		})
	}
}
