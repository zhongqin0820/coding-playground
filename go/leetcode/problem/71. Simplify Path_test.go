package problem

import (
	"fmt"
	"strings"
	"testing"
)

// https://leetcode.com/problems/simplify-path/
func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	queue := []string{}
	for _, str := range strs {
		if strings.Compare(str, "..") == 0 {
			if len(queue) > 0 {
				queue = queue[:len(queue)-1]
			}
		} else if strings.Compare(str, ".") == 0 || strings.Compare(str, "/") == 0 || strings.Compare(str, "") == 0 {
			continue
		} else {
			queue = append(queue, str)
		}
	}
	return "/" + strings.Join(queue, "/")
}

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
		{"/home/../../..", "/"},
	}

	//
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			if res := simplifyPath(ts.input); res != ts.output {
				t.Errorf("expected %s got %s\n", ts.output, res)
			}
		})
	}
}
