package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/word-ladder/submissions/

// Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

// Only one letter can be changed at a time.
// Each transformed word must exist in the word list. Note that beginWord is not a transformed word.

// Note:
// Return 0 if there is no such transformation sequence.
// All words have the same length.
// All words contain only lowercase alphabetic characters.
// You may assume no duplicates in the word list.
// You may assume beginWord and endWord are non-empty and are not the same.

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool, len(wordList))
	for i := 0; i < len(wordList); i++ {
		dict[wordList[i]] = true
	}
	delete(dict, beginWord)
	queue := make([]string, 0, len(wordList))

	var trans func(string) bool
	trans = func(word string) bool {
		bs := []byte(word)
		for i := 0; i < len(bs); i++ {
			diffLetter := bs[i]

			for j := 0; j < 26; j++ {
				b := 'a' + byte(j)
				if b == diffLetter {
					continue
				}

				bs[i] = b
				s := string(bs)
				if dict[s] {
					if s == endWord {
						return true
					}
					queue = append(queue, s)
					delete(dict, s)
				}
			}
			bs[i] = diffLetter
		}
		return false
	}

	// 进行bfs
	queue = append(queue, beginWord)
	for dist := 1; len(queue) > 0; dist++ {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			word := queue[0]
			queue = queue[1:]
			if trans(word) {
				return dist + 1
			}
		}
	}
	return 0
}

func TestLadderLength(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		output    int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		// Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
		// return its length 5.
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
		// Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, ladderLength(ts.beginWord, ts.endWord, ts.wordList))
		})
	}
}
