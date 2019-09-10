package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-product-of-word-lengths/description/
func maxProduct(words []string) int {
	n := len(words)
	masks := make([]int, n)
	for i := 0; i < n; i++ {
		for _, b := range words[i] {
			// 利用比特位来描述words[i]的单词
			// 总共26个字母，对应位置1
			masks[i] |= (1 << uint32(b-'a'))
		}
	}

	var max int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// &操作不为0，则表示这两个单词具有相同的位
			if masks[i]&masks[j] != 0 {
				continue
			}
			// 不相同进行计算
			temp := len(words[i]) * len(words[j])
			if max < temp {
				max = temp
			}
		}
	}
	return max
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		words  []string
		output int
	}{
		{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}, 16},
		{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
		{[]string{"a", "aa", "aaa", "aaaa"}, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxProduct(ts.words))
		})
	}
}
