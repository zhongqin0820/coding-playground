package barrier

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err  error
	Resp string
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)
	responses := make([]barrierResp, requestNumber)
	// loop twice, once for each endpoint
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}
	var hasError bool
	// check to see if there is any error
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			// find an error, we print it prefixed with the word ERROR as we expect in our tests
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}
	// no error found, print every response and the channel will be closed
	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

// accepts a channel to output barrierResp values to and a URL to request
func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	// create an http.Client and set its timeout field
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) *
			time.Millisecond),
	}
	// make the GET call, take the response
	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	// parse it to a byte slice, and send it through the out channel.
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res
}

func captureBarrierOutput(endpoints ...string) string {
	// a pipe allows us to connect an io.Writer interface to an io.Reader interface
	// so that the reader input is the Writer output
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	barrier(endpoints...)
	writer.Close()
	temp := <-out
	return temp
}
