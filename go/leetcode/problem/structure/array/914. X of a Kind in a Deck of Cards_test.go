package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/
func hasGroupsSizeX(deck []int) bool {
	if deck == nil || len(deck) < 1 {
		return false
	}
	// 统计频数
	n := len(deck)
	counter := make(map[int]int, n)
	for _, card := range deck {
		counter[card]++
	}
	// 计算最大公约数
	common := counter[deck[0]]
	for _, count := range counter {
		common = helper914(common, count)
		if common == 1 {
			return false
		}
	}
	return true
}

// 计算最大公约数
func helper914(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func TestHasGroupsSizeX(t *testing.T) {
	tests := []struct {
		deck   []int
		output bool
	}{
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, true},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3}, false},
		{[]int{1}, false},
		{[]int{1, 1}, true},
		{[]int{1, 1, 2, 2, 2, 2}, true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, hasGroupsSizeX(ts.deck))
		})
	}
}
