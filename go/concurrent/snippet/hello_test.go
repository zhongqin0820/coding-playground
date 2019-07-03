package snippet

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestBasicConcurrency(t *testing.T) {
	// control of goroutine
	// 使用父子表示调用关系，不表示关联关系
	// 使用匿名函数，但不阻塞父协程
	t.Run("without sleep", func(t *testing.T) {
		go func(msg string) {
			println(msg)
		}("hello world without sleep")
	})
	// 使用匿名函数，使用time.Sleep阻塞父协程
	t.Run("sleep", func(t *testing.T) {
		go func(msg string) {
			println(msg)
		}("hello world sleep")
		time.Sleep(1e9)
	})

	// 定义一个匿名函数，并将其赋值给变量f
	f := func(msg string) {
		println(msg)
	}
	// 调用多个协程，依旧使用time.Sleep阻塞父协程
	// 问题：当父协程退出时，可能子协程还未执行完毕
	t.Run("multi goroutine sleep", func(t *testing.T) {
		go f("hello")
		go f("end")
		time.Sleep(1e9)
	})

	// 使用sync.WaitGroup对多个协程进行管理
	fw := func(wg *sync.WaitGroup, msg string) {
		defer wg.Done() //有时可以等价于wg.Add(-1)
		println(msg)
	}
	t.Run("multi goroutine waitgroup", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(2)
		go fw(&wg, "hello")
		go fw(&wg, "end")
		wg.Wait()
	})

	// async
	// callback
	// 注意：当函数做参数时应该只传入函数的签名，应当将其看作是一种类型
	callback := func(msg string, f func(s string)) {
		// 调用函数
		go func() {
			f(strings.ToUpper(msg))
		}()
	}

	t.Run("callback", func(t *testing.T) {
		var wg sync.WaitGroup = sync.WaitGroup{}
		wg.Add(1)
		// 注意此处调用回调封装时，传入的是实例，例如string类型的变量实例以及签名为func(string)的函数定义，定义中应该只是对其签名中参数的调用
		callback("hello world callback", func(s string) {
			defer wg.Add(-1) //等价于wg.Done()
			log.Println("call: ", s)
		})
		log.Println("wait for async")
		wg.Wait()
		log.Println("end")
	})
	// callback hell
	// 多层回调
	t.Run("callhell", func(t *testing.T) {
		var wg sync.WaitGroup = sync.WaitGroup{}
		wg.Add(2)
		callback("outter", func(s string) {
			defer wg.Done()
			callback(fmt.Sprintf("inner: %s\n", s), func(s string) {
				defer wg.Done()
				log.Println("call: ", s)
			})
		})
		log.Println("wait for async")
		wg.Wait()
		log.Println("end")
	})

	// sync
	// go test -race
	// sync.Mutex
	// sync.RWMutex
	// sync with mutex
	type Counter struct {
		sync.Mutex //embedding
		value      int
	}
	t.Run("sync with mutex", func(t *testing.T) {
		c := Counter{}
		wg := sync.WaitGroup{}
		for i := 0; i < 4; i++ {
			wg.Add(1)
			go func(i int) {
				c.Lock()
				defer c.Unlock()
				defer wg.Done()
				c.value++
			}(i)
		}
		wg.Wait()
		println(c.value)
	})

	// chan
	t.Run("unbuf chan", func(t *testing.T) {
		c := make(chan string)
		go func() {
			c <- "hello"
			close(c)
		}()
		msg := <-c
		println(msg, "world")
	})

	t.Run("buf chan", func(t *testing.T) {
		c := make(chan int, 2)
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				c <- i
			}
			close(c)
		}()

		go func() {
			defer wg.Done()
			for v := range c {
				println(v)
			}
		}()
		wg.Wait()
	})

	t.Run("direct chan", func(t *testing.T) {
		f := func() <-chan int {
			// 内部定义一定匿名goroutine
			c := make(chan int, 2)
			go func() {
				for i := 0; i < 5; i++ {
					c <- i
				}
				close(c)
			}()
			return c
		}
		for v := range f() {
			println(v)
		}
	})
}
