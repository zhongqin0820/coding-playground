package main

func main() {
	c := make(chan string)
	go func(c chan<- string, msg string) {
		c <- msg
	}(c, "hello world")
	msg := <-c
	println(msg)
}
