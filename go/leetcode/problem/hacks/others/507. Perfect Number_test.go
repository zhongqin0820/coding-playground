package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/perfect-number/
func checkPerfectNumber(num int) bool {
	sum := 1
	root := int(math.Sqrt(float64(num)))
	for i := 2; i <= root; i++ {
		if 0 == num%i {
			sum += (num / i) + i
		}
	}
	return num > 1 && sum == num
}

func TestCheckPerfectNumber(t *testing.T) {
	tests := []struct {
		num    int
		output bool
	}{
		{1, false},
		{2, false},
		{3, false},
		{28, true},
		{120, false},
		{2016, false},
		{32640, false},
		{130816, false},
		{523776, false},
		{2096128, false},
		{8386560, false},
		{100000000, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, checkPerfectNumber(ts.num))
		})
	}
}
