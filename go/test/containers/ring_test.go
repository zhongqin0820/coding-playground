package containers

import (
	"container/ring"
	"fmt"
	"testing"
)

func printRing(r *ring.Ring) {
	r.Do(func(v interface{}) {
		fmt.Print(v.(int), " ")
	})
	fmt.Println()
}

func TestRing(t *testing.T) {
	//创建环形链表
	r := ring.New(5)
	//循环赋值
	for i := 0; i < 5; i++ {
		r.Value = i
		//取得下一个元素
		r = r.Next()
	}
	printRing(r)
	//环的长度
	t.Log(r.Len())

	//移动环的指针
	r.Move(2)

	//从当前指针删除n个元素
	r.Unlink(2)
	printRing(r)

	//连接两个环
	r2 := ring.New(3)
	for i := 0; i < 3; i++ {
		r2.Value = i + 10
		//取得下一个元素
		r2 = r2.Next()
	}
	printRing(r2)

	r.Link(r2)
	printRing(r)
}
