package main

import "time"

// 使用匿名函数，使用time.Sleep阻塞父协程
func main() {
	go func(msg string) {
		println(msg)
	}("hello world sleep")
	time.Sleep(1e9)
}
