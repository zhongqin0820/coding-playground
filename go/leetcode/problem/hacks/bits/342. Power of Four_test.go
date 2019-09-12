package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/power-of-four/
func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}
	// 相当于计算base=4的数
	for num%4 == 0 {
		num /= 4
	}
	return num == 1
}

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		num    int
		output bool
	}{
		{16, true},
		{5, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isPowerOfFour(ts.num))
		})
	}
}
