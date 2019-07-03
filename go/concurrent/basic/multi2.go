package main

import "sync"

// 使用sync.WaitGroup对多个协程进行管理
var fw = func(wg *sync.WaitGroup, msg string) {
	defer wg.Done() //有时可以等价于wg.Add(-1)
	println(msg)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go fw(&wg, "hello")
	go fw(&wg, "end")
	wg.Wait()
}
