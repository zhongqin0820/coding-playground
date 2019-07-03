package main

import "time"

// 定义一个匿名函数，并将其赋值给变量f
var f = func(c chan<- string, msg string) {
	c <- msg
}

func MyPrint(c <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-c:
			println(msg)
		case <-time.After(time.Second * 2):
			// time.After() 返回一个接收通道：<- chan time.Time
			println("break after 2s")
			quitCh <- true
			break
		}
	}
}

func main() {
	c := make(chan string)
	q := make(chan bool)
	go f(c, "hello")
	go MyPrint(c, q)
	go f(c, "world")
	<-q
}
