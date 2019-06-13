package snippet

import (
	"net/http"
	"net/http/httptrace"
	"testing"
)

func TestTrace(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://zyk.ajiao.com", nil)
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			t.Logf("Got conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			t.Logf("DNS info: %+v\n", dnsInfo)
		},
	}
	// replay
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		t.Log(err)
	}
}
