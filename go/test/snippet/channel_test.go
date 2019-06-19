package snippet

import (
	"testing"
	"unsafe"
)

func TestChanSize(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int, 0)
	c3 := make(chan int, 1)
	t.Log("c1:size=", unsafe.Sizeof(c1), "len=", len(c1), "cap=", cap(c1))
	t.Log("c2:size=", unsafe.Sizeof(c2), "len=", len(c2), "cap=", cap(c2))
	t.Log("c3:size=", unsafe.Sizeof(c3), "len=", len(c3), "cap=", cap(c3))
}
