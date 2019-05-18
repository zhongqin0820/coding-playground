package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// string 属于常量,不可以修改,但是可以拼接产生一个新的字符串
	s := "abcdf"
	fmt.Println(s)
	// cannot assign to s[1]
	// s[1] = "1"
	// fmt.Println(s)
	s += "gh"
	fmt.Println(s)
	fmt.Printf("%d \n", unsafe.Sizeof(s)) //16=8+8(8字节的uintptr,8字节的int)

	//array
	var a [5][0]int
	fmt.Printf("%T \n", a)                   //[5][0]int
	fmt.Printf("%#v \n", a)                  //[5][0]int{[0]int{}, [0]int{}, [0]int{}, [0]int{}, [0]int{}}
	fmt.Printf("%d \n", unsafe.Sizeof(a))    //0
	fmt.Printf("%T \n", a[0])                //[0]int
	fmt.Printf("%#v \n", a[0])               //[0]int{}
	fmt.Printf("%d \n", unsafe.Sizeof(a[0])) //0
}
