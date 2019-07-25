package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/palindrome-pairs/

// Given a list of unique words, find all pairs of distinct indices (i, j) in the given list, so that the concatenation of the two words, i.e. words[i] + words[j] is a palindrome.

func palindromePairs(words []string) [][]int {
	n := len(words)
	res := make([][]int, 0, n)
	if n < 2 {
		return res
	}

	hash := make(map[string]int, n)
	for i := 0; i < n; i++ {
		hash[words[i]] = i
	}

	for i := 0; i < n; i++ {
		for k := 0; k <= len(words[i]); k++ {
			left := words[i][:k]
			right := words[i][k:]

			// 如果右侧是回文
			if helper336(right) {
				// 反转该词的左侧，如果能在字典中找到，则加入结果
				leftReverves := helper336Reverse(left)
				if j, ok := hash[leftReverves]; ok && i != j {
					res = append(res, []int{i, j})
				}
			}

			// 如果左侧不是空，且也为回文
			if len(left) != 0 && helper336(left) {
				// 反转该词的右侧，如果能在字典中找到，则加入结果
				rightReverses := helper336Reverse(right)
				if j, ok := hash[rightReverses]; ok && i != j {
					res = append(res, []int{j, i})
				}
			}
		}
	}
	return res
}

// isPalindrome
func helper336(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

func helper336Reverse(s string) string {
	n := len(s)
	bs := []byte(s)
	for i := 0; i < n/2; i++ {
		bs[i], bs[n-i-1] = bs[n-i-1], bs[i]
	}
	return string(bs)
}

func TestPalindromePairs(t *testing.T) {
	tests := []struct {
		words  []string
		output [][]int
	}{
		{[]string{"abcd", "dcba", "lls", "s", "sssll"}, [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}},
		{[]string{"bat", "tab", "cat"}, [][]int{{0, 1}, {1, 0}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, palindromePairs(ts.words))
		})
	}
}
