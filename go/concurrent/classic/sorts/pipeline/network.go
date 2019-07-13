package pipeline

import (
	"bufio"
	"net"
)

func NetworkSource(addr string) <-chan int {
	out := make(chan int, CHANSIZE)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		r := ReaderSource(conn, -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out
}

func NetworkSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		defer listener.Close()

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriterSink(writer, in)
	}()
}
