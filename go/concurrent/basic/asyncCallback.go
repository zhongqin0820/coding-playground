package main

import (
	"log"
	"strings"
	"sync"
)

// callback
// 注意：当函数做参数时应该只传入函数的签名，应当将其看作是一种类型
var callback = func(msg string, f func(s string)) {
	// callback方法开启一个协程并发调度函数：f func(s string)
	go func() {
		f(strings.ToUpper(msg))
	}()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 注意此处调用回调封装时，传入的是实例，例如string类型的变量实例以及签名为func(string)的函数定义，定义中应该只是对其签名中参数的调用
	callback("hello world callback", func(s string) {
		defer wg.Done()
		log.Println("call: ", s)
	})
	log.Println("wait for async call")
	wg.Wait()
	log.Println("end of async call")
}
