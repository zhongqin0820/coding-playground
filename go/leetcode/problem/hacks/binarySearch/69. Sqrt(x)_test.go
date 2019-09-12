package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sqrtx/description/
func mySqrt(x int) int {
	res := x
	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

func naiveMySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, h := 1, x
	for l <= h {
		m := l + (h-l)/2
		sqrt := x / m
		if sqrt == m {
			return m
		} else if sqrt < m { // 说明m*m>x
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return h
}

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x      int
		output int
	}{
		{4, 2},
		{8, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, mySqrt(ts.x))
			ast.Equal(ts.output, naiveMySqrt(ts.x))
		})
	}
}
