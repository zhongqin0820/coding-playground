package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/word-search/

// Given a 2D board and a word, find if the word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	if m*n < len(word) {
		return false
	}

	var dfs func(int, int, int) bool
	dfs = func(r, c, index int) bool {
		if len(word) == index {
			return true
		}
		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != word[index] {
			return false
		}
		temp := board[r][c]
		board[r][c] = 0
		if dfs(r-1, c, index+1) ||
			dfs(r+1, c, index+1) ||
			dfs(r, c-1, index+1) ||
			dfs(r, c+1, index+1) {
			return true
		}
		board[r][c] = temp // backtracking
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func TestExist(t *testing.T) {
	tests := []struct {
		board  [][]byte
		word   string
		output bool
	}{
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "ABCCED", true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "SEE", true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "ABCB", false,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, exist(ts.board, ts.word))
		})
	}
}
