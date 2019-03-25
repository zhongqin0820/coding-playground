package containers

import (
	"container/list"
	"fmt"
)

func IteList() {
	var l = list.New()
	l.PushBack(4)
	e2 := l.PushFront(2)
	l.InsertAfter(3, e2)
	l.InsertBefore(1, e2)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}

// struct is not support to compare
type Person struct {
	Name string
	Age  int
}

func StructFea() {
	p1 := &Person{Name: "A", Age: 12}
	p2 := &Person{Name: "B", Age: 12}
	p3 := &Person{Name: "A", Age: 12}
	if p1 == p2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
	if p1 == p3 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
}
