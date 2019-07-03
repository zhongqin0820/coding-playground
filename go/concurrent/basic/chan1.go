package main

func main() {
	c := make(chan string)
	go func() {
		c <- "hello"
		close(c)
	}()
	msg := <-c
	println(msg, "world")
}
