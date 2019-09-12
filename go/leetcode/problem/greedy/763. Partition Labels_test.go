package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/partition-labels/description/
func partitionLabels(S string) []int {
	maxIndex := [26]int{}
	// 统计对应字符的下标
	for i, b := range S {
		maxIndex[b-'a'] = i
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	begin, end, res := 0, maxIndex[S[0]-'a'], make([]int, 0, len(S))
	for i, b := range S {
		// 在 S[:i+1] 和 S[i:] 中存在相同的字母 S[end]
		// 所以此时不能分隔，仅更新 end
		if i < end {
			end = max(end, maxIndex[b-'a'])
			continue
		}
		// 此时 S[begin:i+1] 中的所有字母都不会出现在其他片段中
		// 可以进行分隔
		res = append(res, i-begin+1)
		// 从 i+1 处作为新片段的起始点
		if begin = i + 1; begin < len(S) {
			end = maxIndex[S[begin]-'a']
		}
	}
	return res
}

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		S      string
		output []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, partitionLabels(ts.S))
		})
	}
}
