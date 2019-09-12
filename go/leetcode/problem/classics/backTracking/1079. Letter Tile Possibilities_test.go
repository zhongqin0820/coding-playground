package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/letter-tile-possibilities/

// You have a set of tiles, where each tile has one letter tiles[i] printed on it.  Return the number of possible non-empty sequences of letters you can make.

// Note:

// 1 <= tiles.length <= 7
// tiles consists of uppercase English letters.

func numTilePossibilities(tiles string) int {
	var count = [26]int{}
	for _, v := range tiles {
		count[v-'A']++
	}

	var dfs func() int

	dfs = func() int {
		sum := 0
		for i := 0; i < 26; i++ {
			if count[i] == 0 {
				continue
			}
			sum++
			count[i]--
			sum += dfs()
			count[i]++
		}
		return sum
	}
	return dfs()
}

func TestNumTilePossibilities(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"AAB", 8},
		// Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
		{"AAABBC", 188},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numTilePossibilities(ts.input))
		})
	}
}
