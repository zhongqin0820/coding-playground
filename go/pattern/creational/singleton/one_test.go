package singleton

import (
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	var globalwg sync.WaitGroup
	globalwg.Add(2)
	// one thread
	t.Run("one", func(t *testing.T) {
		s1 := GetOneInstance()
		s2 := GetOneInstance()
		if s2 == s1 {
			// s2.AddOne()
			t.Logf("instance1=%p,instance2=%p\n", s1, s2)
		} else {
			t.Errorf("instance1=%p,instance2=%p\n", s1, s2)
		}
		s2 = GetOneInstance()
		if s2 == s1 {
			// s2.AddOne()
			t.Logf("instance1=%p,instance2=%p\n", s1, s2)
		} else {
			t.Errorf("instance1=%p,instance2=%p\n", s1, s2)
		}
		t.Logf("value=%d\n", s1.GetValue())
		globalwg.Done()
	})

	// multi thread
	go t.Run("multi", func(t *testing.T) {
		s1 := GetOneInstance()
		var wg sync.WaitGroup
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				s := GetOneInstance()
				if s != s1 {
					t.Errorf("instance1=%p,instance%d=%p\n", s1, i, s)
				} else {
					// s.AddOne()
					t.Log(i)
				}
				time.Sleep(1e9)
			}(i)
		}
		wg.Wait()
		t.Logf("value=%d\n", s1.GetValue())
		globalwg.Done()
	})
	globalwg.Wait()
}
