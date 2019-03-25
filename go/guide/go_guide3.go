package main

import (
	"fmt"
	"sync"
	"time"
)

//
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func exp1() {
	go say("world")
	say("hello")
}

//
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func exp2() {
	a := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

//
func exp3() {
	c := make(chan int, 1)
	c <- 1
	c <- 2
	fmt.Println(<-c)
}

//
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func exp4() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

//
func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func exp5() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	Fibonacci(c, quit)
}

func exp6() {
	tick := time.Tick(10 * time.Millisecond)
	boom := time.After(20 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
		default:
			fmt.Println(" .")
			time.Sleep(5 * time.Millisecond)
		}
	}
}

//检查两个二叉树是否存储了相同的序列的函数
// type Tree struct {
// 	Left *Tree
// 	Value int
// 	Right *Tree
// }
// func Walk(t *tree.Tree, ch chan int) {
// 	sendValue(t, ch)
// 	close(ch)
// }

// func sendValue(t *tree.Tree, ch chan int) {
// 	if t != nil {
// 		sendValue(t.Left, ch)
// 		ch <- t.Value
// 		sendValue(t.Right, ch)
// 	}
// }

// func Same(t1, t2 *tree.Tree) bool {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	go Walk(t1, ch1)
// 	go Walk(t2, ch2)
// 	for i := range ch1 {
// 		if i != <-ch2 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func exp7() {
// 	var ch = make(chan int)
// 	go Walk(tree.New(1), ch)
// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// 	fmt.Println(Same(tree.New(1), tree.New(1)))
// 	fmt.Println(Same(tree.New(1), tree.New(2)))

// }

// 并发爬虫
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var (
	m = make(map[string]int)
	// 互斥锁
	l sync.Mutex
	// Add() Wait() Done()
	i sync.WaitGroup
)

func Crawl(url string, depth int, fetcher Fetcher) {
	defer i.Done()
	if depth <= 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 因为并发的原因，存在很多个子进程同时进行，在子进程中为写操作添加互斥锁
	l.Lock()
	if m[url] == 0 {
		m[url]++
		depth--
		// fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			i.Add(1)
			// goroutine
			go Crawl(u, depth-1, fetcher)
		}
	}
	l.Unlock()
	return
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

//
func exp8() {
	i.Add(1)
	Crawl("http://golang.org/", 4, fetcher)
	i.Wait() // 会一直等待直到子线程任务结束
	for k, _ := range m {
		fmt.Println(k)
	}
	fmt.Println("Over")
}

func main() {
	// exp1()
	// exp2()
	// exp3()
	// exp4()
	// exp5()
	// exp6()
	exp8()
}
