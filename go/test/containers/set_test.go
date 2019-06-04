package containers

import (
	"log"
	"testing"
)

// github.com/deckarep/golang-set
type Bits uint8

const (
	F0 = 1 << iota
	F1
	F2
)

func Set(b, flag Bits) Bits {
	log.Println(b, flag, b|flag)
	return b | flag
}

func Clear(b, flag Bits) Bits {
	log.Println(b, flag, b&^flag)
	return b &^ flag
}

func Toggle(b, flag Bits) Bits {
	log.Println(b, flag, b^flag)
	return b ^ flag
}

func Has(b, flag Bits) bool {
	log.Println(b, flag, (b&flag) != 0)
	return (b & flag) != 0
}

func TestSet(t *testing.T) {
	var b Bits
	b = Set(b, F0)
	b = Toggle(b, F2)
	t.Log([]Bits{F0, F1, F2})
	for i, flag := range []Bits{F0, F1, F2} {
		t.Log(i, Has(b, flag))
	}
}

// github.com/yourbasic/bit
const (
	shift = 6    // log_2 64
	mask  = 0x3f // 即0b00111111
)

// index + position
type Bitset struct {
	data []int64
}

func NewBitSet(n int) *Bitset {
	// 获取位置信息
	index := n >> shift

	set := &Bitset{
		data: make([]int64, index+1),
	}

	// 根据 n 设置 bitset
	set.data[index] |= (1 << uint(n&mask))

	return set
}

func (set *Bitset) Contains(n int) bool {
	// 获取位置信息
	index := n >> shift
	return set.data[index]&(1<<uint(n&mask)) != 0
}

func TestBitset(t *testing.T) {
	set := NewBitSet(65)
	t.Log("set contains 65", set.Contains(65))
	t.Log("set contains 64", set.Contains(64))
}
