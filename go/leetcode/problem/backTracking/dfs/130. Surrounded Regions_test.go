package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/surrounded-regions/description/

// Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.

func solve(board [][]byte) {
	if board == nil || len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	var m, n int = len(board), len(board[0])
	// isEscaped[i][j] == true，表示 (i,j) 点是联通的。
	// 也可以用来查看点 (i,j) 是否已经检查过了
	isEscaped := make([][]bool, m)
	for i := 0; i < m; i++ {
		isEscaped[i] = make([]bool, n)
	}
	// idxM，idxN 分别记录点 (i,j) 的坐标值
	idxM, idxN := make([]int, 0, m*n), make([]int, 0, m*n)

	bfs := func(i, j int) {
		isEscaped[i][j] = true
		idxM = append(idxM, i)
		idxN = append(idxN, j)

		for len(idxM) > 0 {
			i, j := idxM[0], idxN[0]
			idxM, idxN = idxM[1:], idxN[1:]

			// 依次对 (i,j) 的上左下右点进行，bfs检查
			if 0 <= i-1 && board[i-1][j] == 'O' && !isEscaped[i-1][j] {
				idxM = append(idxM, i-1)
				idxN = append(idxN, j)
				isEscaped[i-1][j] = true
			}

			if 0 <= j-1 && board[i][j-1] == 'O' && !isEscaped[i][j-1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j-1)
				isEscaped[i][j-1] = true
			}

			if i+1 < m && board[i+1][j] == 'O' && !isEscaped[i+1][j] {
				idxM = append(idxM, i+1)
				idxN = append(idxN, j)
				isEscaped[i+1][j] = true
			}

			if j+1 < n && board[i][j+1] == 'O' && !isEscaped[i][j+1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j+1)
				isEscaped[i][j+1] = true
			}
		}
	}

	// 联通的区域一定会到达四周，所以从4周开始检查所有的联通区域
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' && !isEscaped[0][j] {
			bfs(0, j)
		}
		if board[m-1][j] == 'O' && !isEscaped[m-1][j] {
			bfs(m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' && !isEscaped[i][0] {
			bfs(i, 0)
		}
		if board[i][n-1] == 'O' && !isEscaped[i][n-1] {
			bfs(i, n-1)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			// 修改内部被占领的 'O'
			if board[i][j] == 'O' && !isEscaped[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		board  [][]byte
		output [][]byte
	}{
		{
			[][]byte{
				[]byte("OXXXXXOO"),
				[]byte("OOOXXXXO"),
				[]byte("XXXXOOOO"),
				[]byte("XOXOOXXX"),
				[]byte("OXOXXXOO"),
				[]byte("OXXOOXXO"),
				[]byte("OXOXXXOO"),
				[]byte("OXXXXOXX"),
			},
			[][]byte{
				[]byte("OXXXXXOO"),
				[]byte("OOOXXXXO"),
				[]byte("XXXXOOOO"),
				[]byte("XXXOOXXX"),
				[]byte("OXXXXXOO"),
				[]byte("OXXXXXXO"),
				[]byte("OXXXXXOO"),
				[]byte("OXXXXOXX"),
			},
		},
		{
			[][]byte{
				[]byte("XXXXOX"),
				[]byte("OXXOOX"),
				[]byte("XOXOOO"),
				[]byte("XOOOXO"),
				[]byte("OOXXOX"),
				[]byte("XOXOXX"),
			},
			[][]byte{
				[]byte("XXXXOX"),
				[]byte("OXXOOX"),
				[]byte("XOXOOO"),
				[]byte("XOOOXO"),
				[]byte("OOXXXX"),
				[]byte("XOXOXX"),
			},
		},
		{
			[][]byte{
				[]byte("OXOOXX"),
				[]byte("OXXXOX"),
				[]byte("XOOXOO"),
				[]byte("XOXXXX"),
				[]byte("OOXOXX"),
				[]byte("XXOOOO"),
			},
			[][]byte{
				[]byte("OXOOXX"),
				[]byte("OXXXOX"),
				[]byte("XOOXOO"),
				[]byte("XOXXXX"),
				[]byte("OOXOXX"),
				[]byte("XXOOOO"),
			},
		},
		{
			[][]byte{
				[]byte("XOXOXOOOXO"),
				[]byte("XOOXXXOOOX"),
				[]byte("OOOOOOOOXX"),
				[]byte("OOOOOOXOOX"),
				[]byte("OOXXOXXOOO"),
				[]byte("XOOXXXOXXO"),
				[]byte("XOXOOXXOXO"),
				[]byte("XXOXXOXOOX"),
				[]byte("OOOOXOXOXO"),
				[]byte("XXOXXXXOOO"),
			},
			[][]byte{
				[]byte("XOXOXOOOXO"),
				[]byte("XOOXXXOOOX"),
				[]byte("OOOOOOOOXX"),
				[]byte("OOOOOOXOOX"),
				[]byte("OOXXOXXOOO"),
				[]byte("XOOXXXXXXO"),
				[]byte("XOXXXXXOXO"),
				[]byte("XXOXXXXOOX"),
				[]byte("OOOOXXXOXO"),
				[]byte("XXOXXXXOOO"),
			},
		},

		{
			[][]byte{
				[]byte("XXXX"),
				[]byte("XOOX"),
				[]byte("XOOX"),
				[]byte("XXXX"),
			},
			[][]byte{
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XXXX"),
			},
			// Explanation:
			// Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
		},
		{
			[][]byte{
				[]byte("OXXOX"),
				[]byte("XOOXO"),
				[]byte("XOXOX"),
				[]byte("OXOOO"),
				[]byte("XXOXO"),
			},
			[][]byte{
				[]byte("OXXOX"),
				[]byte("XXXXO"),
				[]byte("XXXOX"),
				[]byte("OXOOO"),
				[]byte("XXOXO"),
			},
		},
		{
			[][]byte{
				[]byte("XXXX"),
				[]byte("XOOX"),
				[]byte("XXOX"),
				[]byte("XOXX"),
			},
			[][]byte{
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XOXX"),
			},
		},
		{
			[][]byte{
				[]byte("XXXX"),
			},
			[][]byte{
				[]byte("XXXX"),
			},
		},
		{
			[][]byte{
				[]byte("X"),
				[]byte("X"),
				[]byte("X"),
			},
			[][]byte{
				[]byte("X"),
				[]byte("X"),
				[]byte("X"),
			},
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			solve(ts.board)
			ast.Equal(ts.output, ts.board)
		})
	}
}
