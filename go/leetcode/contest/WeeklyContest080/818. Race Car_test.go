package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-80/problems/race-car/
func racecar(target int) int {
	// TODO(zoking)[2019-12-08]: https://www.cnblogs.com/grandyang/p/10360655.html
	return target
}

func TestRacecar(t *testing.T) {
	tests := []struct {
		target int
		output int
	}{
		{3, 2},
		{6, 5},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, racecar(ts.target))
		})
	}
}
