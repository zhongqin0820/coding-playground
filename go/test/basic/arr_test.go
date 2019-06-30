package basic

import "testing"

const LEN = 3

func TestArr(t *testing.T) {
	//声明并初始化
	var a [LEN]int
	// 等同于var a [LEN]int
	var b [LEN]int = [LEN]int{}
	var c [LEN]int = [...]int{1, 2, 9}
	// cannot use [2]int literal (type [2]int) as type [3]int in assignment
	// var c [LEN]int = [...]int{1, 2}
	// cannot use [4]int literal (type [4]int) as type [3]int in assignment
	// var c [LEN]int = [...]int{1, 2, 7, 9}
	var d [LEN]int = [...]int{LEN - 1: 5}
	// var d [LEN]int = [...]int{LEN: 5}
	// cannot make type [3]int
	// var e [LEN]int = make([LEN]int, LEN)
	// 其它变量声明初始化方式
	var e = [LEN]int{}
	f := [LEN]int{}
	t.Logf("\n%v\n%v\n%v\n%v\n", a, b, c, d)
	t.Logf("\n%v\n%v\n", e, f)
}
