package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/iterator-for-combination/
// ref: https://leetcode.com/problems/iterator-for-combination/discuss/451368/C%2B%2B-solution-with-multiple-pointers
type CombinationIterator struct {
	p []int
	c string
	f bool
	l int
	n int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	res := CombinationIterator{
		c: characters,
		f: false,
		l: len(characters),
		n: combinationLength,
		p: make([]int, combinationLength),
	}
	for i := 0; i < combinationLength; i++ {
		res.p[i] = i
	}
	return res
}

func (this *CombinationIterator) Next() string {
	res := make([]byte, 0, this.n)
	if !this.f {
		// ensemble result
		for i := 0; i < this.n; i++ {
			res = append(res, this.c[this.p[i]])
		}
		this.f = true
	} else {
		// update pointers
		for i := this.n - 1; i > -1; i-- {
			if this.p[i] != i+this.l-this.n {
				this.p[i]++
				for j := i + 1; j < this.n; j++ {
					this.p[j] = this.p[i] + j - i
				}
				break
			}
		}
		// ensemble result
		for i := 0; i < this.n; i++ {
			res = append(res, this.c[this.p[i]])
		}
	}
	return string(res)
}

func (this *CombinationIterator) HasNext() bool {
	return this.p[0] != this.l-this.n
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
func TestCombinationIterator(t *testing.T) {
	t.Run("Example 0", func(t *testing.T) {
		ast := assert.New(t)
		obj := Constructor("abc", 2)
		// ["CombinationIterator","next","hasNext","next","hasNext","next","hasNext"]
		// [["abc",2],[],[],[],[],[],[]]
		ast.Equal("ab", obj.Next())
		ast.Equal(true, obj.HasNext())
		ast.Equal("ac", obj.Next())
		ast.Equal(true, obj.HasNext())
		ast.Equal("bc", obj.Next())
		ast.Equal(false, obj.HasNext())
	})
	t.Run("Example 1", func(t *testing.T) {
		ast := assert.New(t)
		obj := Constructor("abc", 3)
		// ["CombinationIterator","next","hasNext","hasNext"]
		// [["abc",3],[],[],[]]
		ast.Equal("abc", obj.Next())
		ast.Equal(false, obj.HasNext())
		ast.Equal(false, obj.HasNext())
	})
	t.Run("Example 2", func(t *testing.T) {
		ast := assert.New(t)
		obj := Constructor("chp", 1)
		// ["CombinationIterator","hasNext","next","hasNext","hasNext","next","next","hasNext","hasNext","hasNext","hasNext"]
		// [["chp",1],[],[],[],[],[],[],[],[],[],[]]
		ast.Equal(true, obj.HasNext())
		ast.Equal("c", obj.Next())
		ast.Equal(true, obj.HasNext())
		ast.Equal("h", obj.Next())
		ast.Equal(true, obj.HasNext())
		ast.Equal("p", obj.Next())
		ast.Equal(false, obj.HasNext())
		ast.Equal(false, obj.HasNext())
		ast.Equal(false, obj.HasNext())
		ast.Equal(false, obj.HasNext())
	})
	t.Run("Example 3", func(t *testing.T) {
		ast := assert.New(t)
		obj := Constructor("gkosu", 3)
		// ["CombinationIterator","hasNext","hasNext","hasNext","next","hasNext","next","next","next","hasNext","next"]
		// [["gkosu",3],[],[],[],[],[],[],[],[],[],[]]
		// [null,true,true,true,"gko",true,"gks","gku","gos",true,"gou"]
		ast.Equal(true, obj.HasNext())
		ast.Equal(true, obj.HasNext())
		ast.Equal(true, obj.HasNext())
		ast.Equal("gko", obj.Next())
		ast.Equal(true, obj.HasNext())
		ast.Equal("gks", obj.Next())
		ast.Equal("gku", obj.Next())
		ast.Equal("gos", obj.Next())
		ast.Equal(true, obj.HasNext())
		ast.Equal("gou", obj.Next())
	})
}
