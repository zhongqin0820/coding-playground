package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/nim-game/

// You are playing the following Nim Game with your friend: There is a heap of stones on the table, each time one of you take turns to remove 1 to 3 stones. The one who removes the last stone will be the winner. You will take the first turn to remove the stones.

// Both of you are very clever and have optimal strategies for the game. Write a function to determine whether you can win the game given the number of stones in the heap.

func canWinNim(n int) bool {
	return n%4 != 0
}

func TestCanWinNim(t *testing.T) {
	tests := []struct {
		n      int
		output bool
	}{
		{4, false},
		// Explanation: If there are 4 stones in the heap, then you will never win the game;
		//            No matter 1, 2, or 3 stones you remove, the last stone will always be
		//            removed by your friend.
		{5, true},
		{41, true},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, canWinNim(ts.n))
		})
	}
}
