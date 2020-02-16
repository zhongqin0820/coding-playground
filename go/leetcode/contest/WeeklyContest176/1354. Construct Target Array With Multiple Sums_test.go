package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/construct-target-array-with-multiple-sums/
// ref: https://leetcode.com/problems/construct-target-array-with-multiple-sums/discuss/510256/JavaC%2B%2BPython-Backtrack-Useless-Test-Cases
func isPossible(target []int) bool {
	sort.Ints(target)
	//
	pre, n, all := 0, len(target), false
	for i := 0; i < n-1; i++ {
		pre += target[i]
		all = target[i] == 1
	}
	if all && target[n-1] == 1 {
		return true
	}
	//
	if target[n-1] <= pre {
		return false
	}
	target[n-1] -= pre
	return isPossible(target)
}

func TestIsPossible(t *testing.T) {
	tests := []struct {
		target []int
		output bool
	}{
		{[]int{9, 3, 5}, true},
		{[]int{1, 1, 1, 2}, false},
		{[]int{8, 5}, true},
		{[]int{1, 1, 2}, false},
		// {[]int{1, 1000000000}, true}, // TLE
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isPossible(ts.target))
		})
	}
}
