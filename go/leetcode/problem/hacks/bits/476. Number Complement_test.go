package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-complement/description/
func findComplement(num int) int {
	temp := num
	res := 0
	//     num=101
	//     res=111
	// res^num=010
	for temp > 0 {
		temp >>= 1
		res <<= 1
		res++
	}
	return res ^ num
}

func TestFindComplement(t *testing.T) {
	tests := []struct {
		num    int
		output int
	}{
		{5, 2},
		{1, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findComplement(ts.num))
		})
	}
}
