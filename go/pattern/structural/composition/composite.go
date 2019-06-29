package composition

import "log"

type Swimmer interface {
	Swim()
}

// SwimmerImpl
type SwimmerImpl struct{}

func NewSwimmerImpl() SwimmerImpl {
	return SwimmerImpl{}
}

func (s SwimmerImpl) Swim() {
	log.Println("SwimmerImpl swim")
}

type Trainer interface {
	Train()
}

// Athlete
type Athlete struct{}

func NewAthlete() Athlete {
	return Athlete{}
}

func (a Athlete) Train() {
	log.Println("Athlete train")
}

// CompositeSwimmerA
type CompositeSwimmerA struct {
	MyAthelete Athlete
	MySwim     func()
}

func NewCompositeSwimmerA(a Athlete, s func()) CompositeSwimmerA {
	return CompositeSwimmerA{MyAthelete: a, MySwim: s}
}

// CompositeSwimmerB
type CompositeSwimmerB struct {
	Swimmer
	Trainer
}

func NewCompositeSwimmerB(si Swimmer, ti Trainer) CompositeSwimmerB {
	return CompositeSwimmerB{Swimmer: si, Trainer: ti}
}

// binary tree
type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func NewTree(v int) *Tree {
	return &Tree{Val: v, Left: nil, Right: nil}
}

func (t *Tree) IterRoot() []int {
	if t.Left == nil && t.Right == nil {
		return []int{t.Val}
	}
	var res []int = []int{t.Val}
	if t.Left != nil {
		res = append(res, t.Left.IterRoot()...)
	}
	if t.Right != nil {
		res = append(res, t.Right.IterRoot()...)
	}
	return res
}
