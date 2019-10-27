package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/design-hashset/
type MyHashSet struct {
	bucket []*val
	length int
}

type val struct {
	v    int
	next *val
}

/** Initialize your data structure here. */
func ConstructorMyHashSet() MyHashSet {
	return MyHashSet{
		bucket: make([]*val, 10000),
		length: 10000,
	}
}

func (this *MyHashSet) Add(key int) {
	if this.bucket[key%this.length] == nil {
		this.bucket[key%this.length] = new(val)
		this.bucket[key%this.length].next = &val{v: key, next: nil}
		return
	}
	// 找到下一个
	pre, cur := this.bucket[key%this.length], this.bucket[key%this.length].next
	for ; cur != nil; pre, cur = cur, cur.next {
		if cur.v == key {
			return
		}
	}
	pre.next = &val{v: key, next: nil}
}

func (this *MyHashSet) Remove(key int) {
	if this.bucket[key%this.length] == nil {
		return
	}
	pre, cur := this.bucket[key%this.length], this.bucket[key%this.length].next
	for ; cur != nil; pre, cur = cur, cur.next {
		if cur.v == key {
			break
		}
	}
	if cur != nil {
		pre.next = cur.next
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if this.bucket[key%this.length] == nil {
		return false
	}
	cur := this.bucket[key%this.length].next
	for ; cur != nil; cur = cur.next {
		if cur.v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func TestMyHashSet(t *testing.T) {
	ast := assert.New(t)
	// MyHashSet hashSet = new MyHashSet();
	obj := ConstructorMyHashSet()
	// hashSet.add(1);
	obj.Add(1)
	// hashSet.add(2);
	obj.Add(2)
	// hashSet.contains(1);    // returns true
	ast.Equal(obj.Contains(1), true)
	// hashSet.contains(3);    // returns false (not found)
	ast.Equal(obj.Contains(3), false)
	// hashSet.add(2);
	obj.Add(2)
	// hashSet.contains(2);    // returns true
	ast.Equal(obj.Contains(2), true)
	// hashSet.remove(2);
	obj.Remove(2)
	// hashSet.contains(2);    // returns false (already removed)
	ast.Equal(obj.Contains(2), false)
}
