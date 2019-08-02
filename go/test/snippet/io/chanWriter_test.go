package snippet

import (
	"log"
	"testing"
)

// ref: https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185
type chanWriter struct {
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{make(chan byte, 1024)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}

func TestChanWriter(t *testing.T) {
	writer := newChanWriter()
	go func() {
		// 写完关闭通道
		defer writer.Close()
		writer.Write([]byte("Stream"))
		writer.Write([]byte("me!"))
	}()

	for c := range writer.Chan() {
		log.Printf("%c", c)
	}
	log.Println()
}
