package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/can-place-flowers/description/
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, size := 0, len(flowerbed); i < size; i++ {
		if flowerbed[i] == 0 &&
			// i==size-1或[i+1]==0
			((i+1 < size && flowerbed[i+1] == 0) || i+1 >= size) &&
			// [i-1]==0或i==0
			((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 < 0) {
			flowerbed[i] = 1
			n--
			if n <= 0 {
				return true
			}
		}
	}
	return n <= 0
}

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		output    bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 1, 0, 1}, 1, true},
		{[]int{0}, 1, true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, canPlaceFlowers(ts.flowerbed, ts.n))
		})
	}
}
