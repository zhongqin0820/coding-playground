package problem

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lru-cache/submissions/
// 解决方案：一个双向列表+Hash Map
type LRUCache struct {
	cache map[int]*list.Element
	l     *list.List
	cap   int
}

// 节点数据
type CacheNode struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]*list.Element),
		l:     list.New(),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// 命中缓存，移动到头部
	if node, ok := this.cache[key]; ok {
		this.l.MoveToFront(node)
		return node.Value.(*CacheNode).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 命中缓存，更新key对应的value并移动到头部
	if node, ok := this.cache[key]; ok {
		node.Value.(*CacheNode).value = value
		this.l.MoveToFront(node)
	} else {
		// 超出缓存最大容量，删除双向列表最后一个节点的数据，并将其从缓存map中删除
		if this.l.Len() == this.cap {
			// 从缓存中删除双向列表中最后一个节点数据
			delete(this.cache, this.l.Back().Value.(*CacheNode).key)
			// 从双向列表中后删除
			this.l.Remove(this.l.Back())
		}
		this.cache[key] = this.l.PushFront(&CacheNode{key, value})
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)
	ast := assert.New(t)
	obj.Put(1, 1)
	obj.Put(2, 2)
	ast.Equal(1, obj.Get(1))
	obj.Put(3, 3)
	ast.Equal(-1, obj.Get(2))
	obj.Put(4, 4)
	ast.Equal(-1, obj.Get(1))
	ast.Equal(3, obj.Get(3))
	ast.Equal(4, obj.Get(4))
}
