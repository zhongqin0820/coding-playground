package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// There are N students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature. For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C. And we defined a friend circle is a group of students who are direct or indirect friends.

// Given a N*N matrix M representing the friend relationship between students in the class. If `M[i][j] = 1`, then the ith and jth students are direct friends with each other, otherwise not. And you have to output the total number of friend circles among all the students.

// Note:
// 1. N is in range [1,200].
// 1. `M[i][i] = 1` for all students.
// 1. If `M[i][j] = 1`, then `M[j][i] = 1`.

// DFS
func findCircleNum(M [][]int) int {
	N, res := len(M), 0
	visited := make([]bool, N)

	var dfs func(int)
	dfs = func(k int) {
		visited[k] = true
		for i := 0; i < N; i++ {
			if M[k][i] == 1 && !visited[i] { //k与i是好友且i没有被访问过，加入当前circle
				dfs(i)
			}
		}
	}

	for i := 0; i < N; i++ {
		if !visited[i] {
			dfs(i)
			res++ //下一个circle
		}
	}

	return res
}

// 并查集来做
func findCircleNumAdv(M [][]int) int {
	N := len(M)
	friend := make([]int, N)
	for i := 0; i < N; i++ {
		friend[i] = i
	}

	union := func(s, d int) {
		for i := range friend {
			if friend[i] == s {
				friend[i] = d
			}
		}
	}

	res := N
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if M[i][j] == 1 { //是好友
				if friend[i] != friend[j] { //还没有联系上的话
					res--
					union(friend[i], friend[j])
				}
			}
		}
	}

	return res
}

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		M      [][]int
		output int
	}{
		{
			[][]int{
				{1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0},
				{1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1},
				{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}},
			3,
		},

		{
			[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			1,
		},

		{
			[][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
			1,
		},

		{
			[][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
			1,
		},

		{
			[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			2,
			// Explanation:The 0th and 1st students are direct friends, so they are in a friend circle. The 2nd student himself is in a friend circle. So return 2.
		},

		{
			[][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}},
			1,
			// Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends, so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findCircleNum(ts.M))
			ast.Equal(ts.output, findCircleNumAdv(ts.M))
		})
	}
}
