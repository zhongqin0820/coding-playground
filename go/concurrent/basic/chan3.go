package main

func main() {
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
}
