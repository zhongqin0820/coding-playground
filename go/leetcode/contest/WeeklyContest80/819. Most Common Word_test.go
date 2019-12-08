package contest

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-80/problems/most-common-word/
func mostCommonWord(paragraph string, banned []string) string {
	// 将paragraph转换为小写，去除非字母字符，按空格切割paragraph得到[]string
	paragraph = strings.ToLower(paragraph)
	ss := strings.FieldsFunc(paragraph, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	// 将banned转换为map保存
	m := make(map[string]struct{}, len(banned))
	for _, str := range banned {
		m[str] = struct{}{}
	}
	// 使用一个map保存paragraph的单词频度
	freqs := make(map[string]int, len(ss))
	res, freq := "", 0
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			if freqs[s]++; freqs[s] > freq {
				freq = freqs[s]
				res = s
			}
		}
	}
	return res
}

func TestMostCommonWord(t *testing.T) {
	tests := []struct {
		paragraph string
		banned    []string
		output    string
	}{
		{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, mostCommonWord(ts.paragraph, ts.banned))
		})
	}
}
