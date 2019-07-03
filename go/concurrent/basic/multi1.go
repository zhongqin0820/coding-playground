package main

import "time"

// 多个goroutine
// 定义一个匿名函数，并将其赋值给变量f
var f = func(msg string) {
	println(msg)
}

func main() {
	go f("hello")
	go f("end")
	time.Sleep(1e9)
}
