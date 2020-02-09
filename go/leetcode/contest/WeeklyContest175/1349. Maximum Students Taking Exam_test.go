package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ref:https://leetcode.com/problems/maximum-students-taking-exam/discuss/503512/C%2B%2BPY3-Dynamic-Programming-with-Bitmap
func maxStudents(seats [][]byte) (ans int) {
	return ans
}

func TestMaxStudents(t *testing.T) {
	tests := []struct {
		seats  [][]byte
		output int
	}{
		{
			[][]byte{
				{'#', '.', '#', '#', '.', '#'},
				{'.', '#', '#', '#', '#', '.'},
				{'#', '.', '#', '#', '.', '#'},
			},
			4,
		},
		{
			[][]byte{
				{'.', '#'},
				{'#', '#'},
				{'#', '.'},
				{'#', '#'},
				{'.', '#'},
			},
			3,
		},
		{
			[][]byte{
				{'#', '.', '.', '.', '#'},
				{'.', '#', '.', '#', '.'},
				{'.', '.', '#', '.', '.'},
				{'.', '#', '.', '#', '.'},
				{'#', '.', '.', '.', '#'},
			},
			10,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxStudents(ts.seats))
		})
	}
}
