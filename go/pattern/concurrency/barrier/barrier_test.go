package barrier

import (
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	// makes two calls to the correct endpoints
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/user-agent",
		}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "User-Agent") {
			t.Fail()
		}
		t.Log(result)
	})
	// have an incorrect endpoint, so it must return an error
	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{
			"http://malformed-url",
			"http://httpbin.org/user-agent",
		}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})
	// return the maximum timeout time so that we can force a timeout error
	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/user-agent",
		}
		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}
