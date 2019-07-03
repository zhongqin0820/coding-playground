package main

// 一个goroutine
// 使用父子表示调用关系，不表示关联关系
// 使用匿名函数，但不阻塞父协程
func main() {
	go func(msg string) {
		println(msg)
	}("hello world without sleep")
}

// 注意，运行的时候，如果使用go run -race则不会退出main.goroutine
// 不使用-race则会退出main.goroutine
