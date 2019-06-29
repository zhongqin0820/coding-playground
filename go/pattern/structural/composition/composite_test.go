package composition

import "testing"

func TestComposite(t *testing.T) {
	// basic
	t.Run("basic", func(t *testing.T) {
		a := NewCompositeSwimmerA(NewAthlete(), NewSwimmerImpl().Swim)
		b := NewCompositeSwimmerB(NewSwimmerImpl(), NewAthlete())
		a.MySwim()
		b.Swim()
		a.MyAthelete.Train()
		b.Train()
	})
	// binary tree
	t.Run("btree", func(t *testing.T) {
		tr := NewTree(1)
		tr.Left = NewTree(2)
		tr.Right = NewTree(5)
		t.Log(tr.IterRoot())
	})
}
