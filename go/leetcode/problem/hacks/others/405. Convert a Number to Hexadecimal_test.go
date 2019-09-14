package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-a-number-to-hexadecimal/description/
func toHex(num int) string {
	var h = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

	if num == 0 {
		return "0"
	}
	var hex string
	for i := 0; i < 8 && num != 0; i++ {
		hex = h[num&15] + hex
		num >>= 4 //等价于num/=16
	}
	return hex
}

func TestToHex(t *testing.T) {
	tests := []struct {
		num    int
		output string
	}{
		{26, "1a"},
		{-1, "ffffffff"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, toHex(ts.num))
		})
	}
}
