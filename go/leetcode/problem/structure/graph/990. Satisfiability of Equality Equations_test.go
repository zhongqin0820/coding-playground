package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/satisfiability-of-equality-equations/
func equationsPossible(equations []string) bool {
	// 根据等式构建并查集
	u := newUnion(26)
	for _, e := range equations {
		if e[1] == '=' {
			u.unite(e[0]-'a', e[3]-'a')
		}
	}
	// 根据不等式，判断它们是否有parent
	for _, e := range equations {
		if e[1] == '!' && u.find(e[0]-'a') == u.find(e[3]-'a') {
			return false
		}
	}

	return true
}

// 并查集合
type union struct {
	parent []byte
}

func newUnion(size int) *union {
	parent := make([]byte, size)
	for i := range parent {
		parent[i] = byte(i)
	}
	return &union{
		parent: parent,
	}
}

func (u *union) find(i byte) byte {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}

func (u *union) unite(x, y byte) {
	u.parent[u.find(x)] = u.find(y)
}

func TestEquationsPossible(t *testing.T) {
	tests := []struct {
		equations []string
		output    bool
	}{
		{[]string{"a==b", "b!=a"}, false},
		{[]string{"b==a", "a==b"}, true},
		{[]string{"a==b", "b==c", "a==c"}, true},
		{[]string{"a==b", "b!=c", "c==a"}, false},
		{[]string{"c==c", "b==d", "x!=z"}, true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, equationsPossible(ts.equations))
		})
	}
}
