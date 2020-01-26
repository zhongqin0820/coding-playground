package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/jump-game/
// solution 1: backtrack
func canJump(nums []int) bool {
	return helper55(nums, 0)
}

func helper55(nums []int, pos int) bool {
	if pos == len(nums)-1 {
		return true
	}
	// cal the maximum steps can jump
	steps := pos + nums[pos]
	if steps > len(nums)-1 {
		steps = len(nums) - 1
	}
	// backtracking
	for i := pos + 1; i <= steps; i++ {
		if helper55(nums, i) {
			return true
		}
	}

	return false
}

// solution 2: memo
type Index int

const (
	GOOD Index = iota
	BAD
	UNKNOWN
)

func canJumpII(nums []int) bool {
	var memo = make([]Index, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = UNKNOWN
	}
	memo[len(nums)-1] = GOOD
	return helper55II(nums, memo, 0)
}

func helper55II(nums []int, memo []Index, pos int) bool {
	if memo[pos] != UNKNOWN {
		return memo[pos] == GOOD
	}
	steps := pos + nums[pos]
	if steps > len(nums)-1 {
		steps = len(nums) - 1
	}
	for i := pos + 1; i <= steps; i++ {
		if helper55II(nums, memo, i) {
			memo[i] = GOOD
			return true
		}
	}
	memo[pos] = BAD
	return false
}

// solution 3: dynamic programming bottom-up
func canJumpIII(nums []int) bool {
	// init
	memo := make([]Index, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = UNKNOWN
	}
	memo[len(nums)-1] = GOOD
	// bottom-up
	for i := len(nums) - 2; i >= 0; i-- {
		// find maximum steps in index i
		steps := i + nums[i]
		if steps > len(nums)-1 {
			steps = len(nums) - 1
		}
		// bottom-up result
		for j := i + 1; j <= steps; j++ {
			if memo[j] == GOOD {
				memo[i] = GOOD
				break
			}
		}

	}
	// check final
	return memo[0] == GOOD
}

// solution 4 : greedy
func canJumpIV(nums []int) bool {
	pos := len(nums) - 1
	for i := pos; i >= 0; i-- {
		// 从后向前，如果i+nums[i]>=pos说明可通过i达到pos
		if i+nums[i] >= pos {
			pos = i
		}
	}
	return pos == 0
}

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums   []int
		output bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, canJump(ts.nums))
			ast.Equal(ts.output, canJumpII(ts.nums))
			ast.Equal(ts.output, canJumpIII(ts.nums))
			ast.Equal(ts.output, canJumpIV(ts.nums))
		})
	}
}
