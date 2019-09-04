package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func solveSudoku(board [][]byte) {
	helper37(board, 0)
}

func helper37(board [][]byte, k int) bool {
	if k == 81 {
		return true
	}
	r, c := k/9, k%9
	if board[r][c] != '.' {
		return helper37(board, k+1)
	}

	/* bi, bj 是 rc 所在块的左上角元素的索引值 */
	bi, bj := r/3*3, c/3*3
	// 按照数独的规则，检查 b 能否放在 board[r][c]
	isValid := func(b byte) bool {
		for n := 0; n < 9; n++ {
			if board[r][n] == b || board[n][c] == b || board[bi+n/3][bj+n%3] == b {
				return false
			}
		}
		return true
	}

	for b := byte('1'); b <= '9'; b++ {
		if isValid(b) {
			board[r][c] = b
			if helper37(board, k+1) {
				return true
			}
		}
	}
	board[r][c] = '.'

	return false
}

func TestSolveSudoku(t *testing.T) {
	tests := []struct {
		board  [][]byte
		output [][]byte
	}{
		{
			[][]byte{
				[]byte("..9748..."),
				[]byte("7........"),
				[]byte(".2.1.9..."),
				[]byte("..7...24."),
				[]byte(".64.1.59."),
				[]byte(".98...3.."),
				[]byte("...8.3.2."),
				[]byte("........6"),
				[]byte("...2759.."),
			},
			[][]byte{
				[]byte("519748632"),
				[]byte("783652419"),
				[]byte("426139875"),
				[]byte("357986241"),
				[]byte("264317598"),
				[]byte("198524367"),
				[]byte("975863124"),
				[]byte("832491756"),
				[]byte("641275983"),
			},
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			solveSudoku(ts.board)
			ast.Equal(ts.output, ts.board)
		})
	}
}
