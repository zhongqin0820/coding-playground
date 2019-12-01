package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
// ref: https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/discuss/441422/c++-0ms-short-and-simple
// 每个玩家都有8个解 = 3 + 3 + 2
// 3行，3列，2对角线
// 每个玩家按顺序放置，奇数A，偶数B；谁先获得解谁赢
func tictactoe(moves [][]int) string {
	// 每个玩家8个可行解
	A, B := make([]int, 8, 8), make([]int, 8, 8)
	var player []int
	size := len(moves)
	for i := 0; i < size; i++ {
		r, c := moves[i][0], moves[i][1]
		if i%2 == 0 {
			player = A
		} else {
			player = B
		}
		player[r]++   // 0,1,2为行解
		player[3+c]++ // 3,4,5为列解
		if r == c {   // 6为正对角线
			player[6]++
		}
		if r == 2-c { // 7为反对角线
			player[7]++
		}
	}

	for i := 0; i < 8; i++ {
		if A[i] == 3 {
			return "A"
		}
		if B[i] == 3 {
			return "B"
		}
	}

	if size == 9 { // 无解时，若全走完步数，打印Draw
		return "Draw"
	}
	return "Pending"
}

func TestTictactoe(t *testing.T) {
	tests := []struct {
		moves  [][]int
		output string
	}{
		{[][]int{[]int{0, 0}, []int{2, 0}, []int{1, 1}, []int{2, 1}, []int{2, 2}}, "A"},
		{[][]int{[]int{0, 0}, []int{1, 1}, []int{0, 1}, []int{0, 2}, []int{1, 0}, []int{2, 0}}, "B"},
		{[][]int{[]int{0, 0}, []int{1, 1}, []int{2, 0}, []int{1, 0}, []int{1, 2}, []int{2, 1}, []int{0, 1}, []int{0, 2}, []int{2, 2}}, "Draw"},
		{[][]int{[]int{0, 0}, []int{1, 1}}, "Pending"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, tictactoe(ts.moves))
		})
	}
}
