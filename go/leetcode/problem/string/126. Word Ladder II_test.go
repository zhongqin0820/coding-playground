package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/word-ladder-ii/

// Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:
// Only one letter can be changed at a time
// Each transformed word must exist in the word list. Note that beginWord is not a transformed word.

// Note:
// Return an empty list if there is no such transformation sequence.
// All words have the same length.
// All words contain only lowercase alphabetic characters.
// You may assume no duplicates in the word list.
// You may assume beginWord and endWord are non-empty and are not the same.

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := map[string]bool{}
	for _, w := range wordList {
		dict[w] = true
	}
	if !dict[endWord] {
		return [][]string{}
	}
	// m 是以当前word为开始可以转化的word的字符列表
	m := map[string][]string{}
	src, dst := map[string]bool{}, map[string]bool{}
	src[beginWord], dst[endWord] = true, true
	b, bw := false, false
	for 0 != len(src) && 0 != len(dst) && !b {
		if len(src) > len(dst) {
			src, dst, bw = dst, src, !bw
		}
		for w, _ := range src {
			dict[w] = false
		}
		src1 := map[string]bool{}
		for w, _ := range src {
			bs := []byte(w)
			for i := 0; i < len(bs); i++ {
				for j := byte('a'); j <= 'z'; j++ {
					bs[i] = j
					s, t := w, string(bs)
					if dst[t] {
						if bw {
							s, t = t, s
						}
						m[t], b = append(m[t], s), true
					} else if dict[t] {
						src1[t] = true
						if bw {
							s, t = t, s
						}
						m[t] = append(m[t], s)
					}
				}
				bs[i] = w[i]
			}
		}
		src = src1
	}

	// 回溯法
	res, ar := [][]string{}, []string{endWord}
	var cal func(string)
	cal = func(s string) {
		l := len(ar)
		if s == beginWord {
			ar1 := make([]string, l)
			copy(ar1, ar)
			for i := 0; i < l>>1; i++ {
				ar1[i], ar1[l-i-1] = ar1[l-i-1], ar1[i]
			}
			res = append(res, ar1)
			return
		}
		for _, t := range m[s] {
			ar = append(ar, t)
			cal(t)
			ar = ar[:l]
		}
	}

	cal(endWord)
	return res
}

func TestFindLadders(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		output    [][]string
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, [][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}}},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, [][]string{}},
		// Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findLadders(ts.beginWord, ts.endWord, ts.wordList))
		})
	}
}
