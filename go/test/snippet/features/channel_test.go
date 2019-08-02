package snippet

import (
	"log"
	"testing"
	"unsafe"
)

// test the info of buffered chan & unbuffered chan
func TestChanSize(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int, 0)
	c3 := make(chan int, 1)
	// c2 <- 1 // put 1 elements will caused c1 blocked
	// c1 <- 1 // put 1 elements will caused c2 blocked
	c3 <- 1
	// c3 <- 2 // put 2 elements will caused c3 blocked
	t.Log("c1:size=", unsafe.Sizeof(c1), "len=", len(c1), "cap=", cap(c1))
	t.Log("c2:size=", unsafe.Sizeof(c2), "len=", len(c2), "cap=", cap(c2))
	t.Log("c3:size=", unsafe.Sizeof(c3), "len=", len(c3), "cap=", cap(c3))
}

func TestRangeUsages(t *testing.T) {
	a := []int{1, 10}
	var b map[int]string = map[int]string{1: "a", 2: "b"}
	c := make(chan int, 2)
	c <- 1
	c <- 4
	close(c) //close chan in case of range
	t.Log("-----------------")
	for i := range a {
		t.Logf("%d", i)
	}
	for _, v := range a {
		t.Logf("%d", v)
	}
	for i, v := range a {
		t.Logf("%d,%d", i, v)
	}
	t.Log("-----------------")
	for k, v := range b {
		t.Logf("%d,%s", k, v)
	}
	for k := range b {
		t.Logf("%d", k)
	}
	for _, v := range b {
		t.Logf("%s", v)
	}
	t.Log("-----------------")
	for v := range c {
		t.Logf("%d", v)
	}
	t.Log("-----------------")
}

func TestCloseChan(t *testing.T) {
	a := make(chan int, 2)
	go func() {
		for i := 0; i < 3; i++ {
			a <- i
		}
		close(a)
		// close(a) // close the chan 2 times will panic
		// panic: close of closed channel
	}()

	for v := range a {
		log.Println(v)
	}
}
