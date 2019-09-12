package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Implement a trie with insert, search, and startsWith methods.

// Note:
// You may assume that all inputs are consist of lowercase letters a-z.
// All inputs are guaranteed to be non-empty strings.

type Trie struct {
	val byte
	// 每个字母开始的子字典树
	sons [26]*Trie
	end  int
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	n := len(word)
	for i := 0; i < n; i++ {
		// 找到对应子树
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie{val: word[i]}
		}
		// 当前节点=字典子树
		node = node.sons[idx]
	}
	node.end++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	n := len(word)
	for i := 0; i < n; i++ {
		idx := word[i] - 'a'
		// 字典树找不到匹配
		if node.sons[idx] == nil {
			return false
		}
		// 当前节点=字典子树
		node = node.sons[idx]
	}
	if node.end > 0 {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	n := len(prefix)
	for i := 0; i < n; i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func TestImplementTrie(t *testing.T) {
	ast := assert.New(t)
	trie := ConstructorTrie()
	// Trie trie = new Trie();
	trie.Insert("apple")
	// trie.insert("apple");
	ast.Equal(true, trie.Search("apple"))
	// trie.search("apple");   // returns true
	ast.Equal(false, trie.Search("app"))
	// trie.search("app");     // returns false
	ast.Equal(true, trie.StartsWith("app"))
	// trie.startsWith("app"); // returns true
	trie.Insert("app")
	// trie.insert("app");
	ast.Equal(true, trie.Search("app"))
	// trie.search("app");     // returns true
}
