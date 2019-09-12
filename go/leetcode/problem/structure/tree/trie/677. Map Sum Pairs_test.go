package problem

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/map-sum-pairs/description/

// Implement a MapSum class with insert, and sum methods.
// For the method insert, you'll be given a pair of (string, integer). The string represents the key and the integer represents the value. If the key already existed, then the original key-value pair will be overridden to the new one.
// For the method sum, you'll be given a string representing the prefix, and you need to return the sum of all the pairs' value whose key starts with the prefix.

type MapSum struct {
	m    map[string]int
	keys []string
}

/** Initialize your data structure here. */
func ConstructorMapSum() MapSum {
	return MapSum{
		m:    make(map[string]int, 128),
		keys: make([]string, 0, 128),
	}
}

func (ms *MapSum) Insert(key string, val int) {
	ms.m[key] = val

	i := sort.SearchStrings(ms.keys, key)
	// 遇到重复的 key ，不操作
	if i < len(ms.keys) && ms.keys[i] == key {
		return
	}

	ms.keys = append(ms.keys, "")
	copy(ms.keys[i+1:], ms.keys[i:])
	ms.keys[i] = key
}

func (ms *MapSum) Sum(prefix string) int {
	res := 0
	i := sort.SearchStrings(ms.keys, prefix)
	for i < len(ms.keys) && strings.HasPrefix(ms.keys[i], prefix) {
		res += ms.m[ms.keys[i]]
		i++
	}
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
func TestMapSumPairs(t *testing.T) {
	ast := assert.New(t)
	obj := ConstructorMapSum()
	obj.Insert("apple", 3)
	// Input: insert("apple", 3), Output: Null
	ast.Equal(3, obj.Sum("ap"))
	// Input: sum("ap"), Output: 3
	obj.Insert("app", 2)
	// Input: insert("app", 2), Output: Null
	ast.Equal(5, obj.Sum("ap"))
	// Input: sum("ap"), Output: 5
}
