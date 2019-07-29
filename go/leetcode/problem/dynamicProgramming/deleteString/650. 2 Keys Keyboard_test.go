package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/2-keys-keyboard/

// Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:

// Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).

// Paste: You can paste the characters which are copied last time.

// Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get n 'A'.

// Note:
// The n will be in the range [1, 1000].
func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return i + minSteps(n/i)
		}
	}

	return n
}

func TestMinSteps(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{3, 3},
		// Intitally, we have one character 'A'.
		// In step 1, we use Copy All operation.
		// In step 2, we use Paste operation to get 'AA'.
		// In step 3, we use Paste operation to get 'AAA'.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minSteps(ts.n))
		})
	}
}
