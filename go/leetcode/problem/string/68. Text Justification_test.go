package problem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/text-justification/

// Given an array of words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.
// You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
// Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.
// For the last line of text, it should be left justified and no extra space is inserted between words.

// Note:
// A word is defined as a character sequence consisting of non-space characters only.
// Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
// The input array words contains at least one word.

func fullJustify(words []string, maxWidth int) []string {
	results := make([]string, 0)
	count, start := 0, 0
	for i, w := range words {
		if count == 0 {
			count, start = len(w), i
		} else {
			if count+1+len(w) <= maxWidth {
				count += 1 + len(w)
			} else {
				v := helper68(count, words[start:i], false, maxWidth)
				results = append(results, v)
				// new start
				count, start = len(w), i
			}
		}
	}

	v := helper68(count, words[start:], true, maxWidth)
	results = append(results, v)
	return results
}

func helper68(count int, words []string, last bool, maxWidth int) string {
	if last {
		return strings.Join(words, " ") + strings.Repeat(" ", maxWidth-count)
	}

	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-count)
	}

	// len(words) >= 2
	extraSpace := maxWidth - count
	gap := len(words) - 1
	space := extraSpace / gap
	leftCount := extraSpace % gap

	result := ""
	if leftCount > 0 {
		// gap + extra_for_left + origin space
		join := strings.Repeat(" ", space+1+1)
		result = strings.Join(words[0:(leftCount+1)], join)
		words = words[leftCount+1:]
	}

	join := strings.Repeat(" ", space+1)
	if result != "" {
		result = result + join + strings.Join(words, join)
	} else {
		result = strings.Join(words, join)
	}

	return result
}

func TestFullJustify(t *testing.T) {
	tests := []struct {
		words    []string
		maxWidth int
		output   []string
	}{
		{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16, []string{"This    is    an", "example  of text", "justification.  "}},
		{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16, []string{"What   must   be", "acknowledgment  ", "shall be        "}},
		// Explanation: Note that the last line is "shall be    " instead of "shall     be",
		//            because the last line must be left-justified instead of fully-justified.
		//            Note that the second line is also left-justified becase it contains only one word.
		{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20, []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, fullJustify(ts.words, ts.maxWidth))
		})
	}
}
