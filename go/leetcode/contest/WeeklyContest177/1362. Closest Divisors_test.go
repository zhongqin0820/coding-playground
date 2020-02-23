package contest

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/closest-divisors/
func closestDivisors(num int) []int {
	// 中间开始，双指针
	i := int(math.Sqrt(float64(num + 2)))
	j := i
	for temp := i * j; i >= 1 && j <= num; temp = i * j {
		if temp == num+1 || temp == num+2 {
			return []int{i, j}
		}
		if temp > num+2 {
			i--
		} else if temp < num+1 {
			j++
		}
	}
	return []int{i, j}
}

func TestClosestDivisors(t *testing.T) {
	tests := []struct {
		num    int
		output []int
	}{
		{2, []int{2, 2}},
		{999, []int{25, 40}},
		{123, []int{5, 25}},
		{8, []int{3, 3}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, closestDivisors(ts.num))
		})
	}
}
