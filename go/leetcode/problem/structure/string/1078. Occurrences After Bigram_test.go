package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/occurrences-after-bigram/
func findOcurrences(t string, a string, b string) []string {
	res := make([]string, 0, 1024)
	n := len(t)
	for i := range t {
		sep, sec, end := i+len(a), i+len(a)+1, i+len(a)+1+len(b)
		if sep >= n || sec >= n || end >= n {
			break
		}
		if (i > 0 && t[i-1] != ' ') || t[sep] != ' ' || t[end] != ' ' {
			continue
		}
		if t[i] != a[0] && t[sec] != b[0] {
			continue
		}
		if match1078(t[i:sep], a) && match1078(t[sec:end], b) {
			for j := end + 1; j < n; j++ {
				if t[j] == ' ' {
					res = append(res, t[end+1:j])
					break
				}
				if j == n-1 {
					res = append(res, t[end+1:j+1])
					break
				}
			}
		}
	}
	return res
}

func match1078(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFindOcurrences(t *testing.T) {
	tests := []struct {
		t      string
		a      string
		b      string
		output []string
	}{
		{"alice is a good girl she is a good student", "a", "good", []string{"girl", "student"}},
		{"we will we will rock you", "we", "will", []string{"we", "rock"}},
		{"alice is aa good girl she is a good student", "a", "good", []string{"student"}},
		{"ypkk lnlqhmaohv lnlqhmaohv lnlqhmaohv ypkk ypkk ypkk ypkk ypkk ypkk lnlqhmaohv lnlqhmaohv lnlqhmaohv lnlqhmaohv ypkk ypkk ypkk lnlqhmaohv lnlqhmaohv ypkk", "lnlqhmaohv", "ypkk", []string{"ypkk", "ypkk"}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findOcurrences(ts.t, ts.a, ts.b))
		})
	}
}
