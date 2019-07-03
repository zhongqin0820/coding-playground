package basic

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
	"time"
)

// demo code from https://golang.org/pkg/sync/#Pool
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	//
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func TestLog(t *testing.T) {
	Log(os.Stdout, "path", "/search?q=flowers")
}
