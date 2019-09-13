package problem

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := list.New()
	for i := range s {
		switch {
		case s[i] == '(' || s[i] == '[' || s[i] == '{':
			stack.PushBack(s[i])
		case s[i] == ')' || s[i] == ']' || s[i] == '}':
			if stack.Len() == 0 {
				return false
			}
			left := stack.Remove(stack.Back()).(byte)
			switch {
			case left == '(' && s[i] == ')':
				continue
			case left == '[' && s[i] == ']':
				continue
			case left == '{' && s[i] == '}':
				continue
			default:
				return false
			}
		default:
			return false
		}
	}
	return true && stack.Len() == 0
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		s      string
		output bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"[", false},
		{"((", false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isValid(ts.s))
		})
	}
}
