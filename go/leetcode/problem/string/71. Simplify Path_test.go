package problem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/simplify-path/

// Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.

// In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level. For more information, see: Absolute path vs relative path in Linux/Unix

// Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) must not end with a trailing /. Also, the canonical path must be the shortest string representing the absolute path.

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
		// Explanation: Note that there is no trailing slash after the last directory name.
		{"/../", "/"},
		// Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
		{"/home//foo/", "/home/foo"},
		// Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
		{"/home/../../..", "/"},
	}

	//
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, simplifyPath(ts.input))
		})
	}
}
