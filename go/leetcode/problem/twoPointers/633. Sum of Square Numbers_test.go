package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-square-numbers/description/
func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		temp := i*i + j*j
		if temp == c {
			return true
		} else if temp < c {
			i++
		} else {
			j--
		}
	}
	return false
}

func TestJudgeSquareSum(t *testing.T) {
	tests := []struct {
		c      int
		output bool
	}{
		{5, true},
		{4, true},
		{3, false},
		{2, true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, judgeSquareSum(ts.c))
		})
	}
}
