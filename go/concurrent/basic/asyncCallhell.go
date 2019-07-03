package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

// callback hell
// 注意：当函数做参数时应该只传入函数的签名，应当将其看作是一种类型
var callback = func(msg string, f func(s string)) {
	// 调用函数
	go func() {
		f(strings.ToUpper(msg))
	}()
}

// callback hell
// 多层回调
func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(2)
	// 此处是回调函数的调用
	callback("outter", func(s string) {
		defer wg.Done()
		callback(fmt.Sprintf("inner: %s", s), func(s string) {
			defer wg.Done()
			log.Println("call: ", s)
		})
	})
	log.Println("wait for async callback")
	wg.Wait()
	log.Println("end of async call")
}
