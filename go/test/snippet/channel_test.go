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

func TestRange(t *testing.T) {
	a := []int{1, 10}
	var b map[int]string = map[int]string{1: "a", 2: "b"}
	c := make(chan int, 2)
	c <- 1
	c <- 4
	close(c) //close chan in case of range
	for i, v := range a {
		t.Logf("%d,%d\t", i, v)
	}
	t.Log()
	for k, v := range b {
		t.Logf("%d,%s\t", k, v)
	}
	t.Log()
	for v := range c {
		t.Logf("%d\t", v)
	}
}
