package problem

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-bits/description/
func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}

func reverseBitsII(n uint32) uint32 {
	// for 8 bit binary number abcdefgh, the process is as follow:
	// abcdefgh -> efghabcd -> ghefcdab -> hgfedcba
	n = (n >> 16) | (n << 16)
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)
	return n
}

func TestReverseBits(t *testing.T) {
	tests := []struct {
		num    uint32
		output uint32
	}{
		{uint32(43261596), uint32(964176192)},
		{uint32(4294967293), uint32(3221225471)},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, reverseBits(ts.num))
			ast.Equal(ts.output, reverseBitsII(ts.num))
		})
	}
}
