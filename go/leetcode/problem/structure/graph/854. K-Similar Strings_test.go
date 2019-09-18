package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/k-similar-strings/
func kSimilarity(a string, b string) int {
	if a == b {
		return 0
	}
	visited := make(map[string]bool, 4096)
	visited[a] = true
	queue := make([]string, 1, 4096)
	queue[0] = a
	res := 0
	strSize := len(a)

	for {
		res++
		for countDown := len(queue); countDown > 0; countDown-- {
			s := queue[0]
			queue = queue[1:]
			i := 0
			for ; s[i] == b[i]; i++ {
			}

			for j := i + 1; j < strSize; j++ {
				// 跳过相同字符，当需要交换的s[i]与s[j]相同不交换
				if s[j] == b[j] || s[i] != b[j] {
					continue
				}
				// 交换后i与j后的字符串
				temp := helper854(s, i, j)
				if temp == b {
					return res
				}
				if !visited[temp] {
					visited[temp] = true
					queue = append(queue, temp)
				}
			}

		}
	}
}

func helper854(s string, i, j int) string {
	bs := []byte(s)
	bs[i], bs[j] = bs[j], bs[i]
	return string(bs)
}

func TestKSimilarity(t *testing.T) {
	tests := []struct {
		A      string
		B      string
		output int
	}{
		{"ab", "ba", 1},
		{"abc", "bca", 2},
		{"abac", "baca", 2},
		{"aabc", "abca", 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, kSimilarity(ts.A, ts.B))
		})
	}
}
