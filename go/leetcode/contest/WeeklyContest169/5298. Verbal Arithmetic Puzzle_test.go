package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-169/problems/verbal-arithmetic-puzzle/
func isSolvable(words []string, result string) bool {
	// TODO[zoking](2019-12-29): https://leetcode.com/problems/verbal-arithmetic-puzzle/discuss/463953/Java-Backtracking-Clean-Code-and-Optimized-700ms
	return false
}

func TestIsSolvable(t *testing.T) {
	tests := []struct {
		words  []string
		result string
		output bool
	}{
		{[]string{"SEND", "MORE"}, "MONEY", true},
		{[]string{"SIX", "SEVEN", "SEVEN"}, "TWENTY", true},
		{[]string{"THIS", "IS", "TOO"}, "FUNNY", true},
		{[]string{"LEET", "CODE"}, "POINT", false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isSolvable(ts.words, ts.result))
		})
	}
}
