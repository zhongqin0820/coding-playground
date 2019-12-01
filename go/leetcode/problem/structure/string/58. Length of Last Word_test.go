package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLastWord(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	res := 0
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 { // 第二次遇到' '
				return res
			}
			continue // 第一次遇到' '
		}
		res++
	}
	return res
}

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"Hello World", 5},
		{"a ", 1},
		{"    ", 0},
		{" a", 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, lengthOfLastWord(ts.s))
		})
	}
}
