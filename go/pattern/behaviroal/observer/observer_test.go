package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	s1 := NewSub(1)
	s2 := NewSub(2)
	s3 := NewSub(3)
	p := NewPub()
	//
	t.Run("AddSub", func(t *testing.T) {
		p.Register(s1)
		p.Register(s2)
		p.Register(s3)
		if len(p.subList) != 3 {
			t.Fail()
		}
	})
	//
	t.Run("DeSub", func(t *testing.T) {
		p.Deregister(s2)
		if len(p.subList) != 2 {
			t.Errorf("expected 3, get %d\n", p.subList)
		}
		//
		for s, _ := range p.subList {
			c, ok := s.(*sub)
			if !ok {
				t.Fail()
			}
			if c.id == 2 {
				t.Fail()
			}
		}
	})
	//
	t.Run("Notify", func(t *testing.T) {
		if len(p.subList) == 0 {
			t.Errorf("publist is empty")
		}
		for s, _ := range p.subList {
			c, ok := s.(*sub)
			if !ok {
				t.Fail()
			}
			if c.theme != "default" {
				t.Errorf("expected default, get %s\n", c.theme)
			}
			// t.Log(c.id, c.theme)
		}
		// update theme
		var newTheme Theme = "hello"
		p.Notify(newTheme)
		for s, _ := range p.subList {
			c, ok := s.(*sub)
			if !ok {
				t.Fail()
			}
			if c.theme != newTheme {
				t.Errorf("expected %s, get %s\n", newTheme, c.theme)
			}
			// t.Log(c.id, c.theme)
		}
	})
}
