package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-burgers-with-no-waste-of-ingredients/
func numOfBurgers(x int, y int) []int {
	// a*4+b*2 = x ...(1)
	// a + b = y ===> b = y-a ...(2)
	// 2*a + 2*y = x ...将(2)代入(1)中
	// 得: a = (x-2y)/2
	a := (x - 2*y) / 2
	b := y - a
	if a < 0 || b < 0 || a*4+b*2 != x {
		return []int{}
	}
	return []int{a, b}
}

func TestNumOfBurgers(t *testing.T) {
	tests := []struct {
		x, y   int
		output []int
	}{
		{16, 7, []int{1, 6}},
		{17, 4, []int{}},
		{4, 17, []int{}},
		{0, 0, []int{0, 0}},
		{2, 1, []int{0, 1}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numOfBurgers(ts.x, ts.y))
		})
	}
}
