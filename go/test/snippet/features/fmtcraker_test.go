package snippet

import (
	"fmt"
	"testing"
)

type Str string

func (s Str) String() string {
	// Sprintf format %s with arg s causes recursive String method call
	// return fmt.Sprintf("Str: %s", s)
	return fmt.Sprintf("Str: %s", string(s))
}

// TestString is a snippet from https://blog.csdn.net/CodyGuo/article/details/51418286
func TestString(t *testing.T) {
	var input, expected Str = "hi", "hi"
	t.Run("crack", func(t *testing.T) {
		if input != expected {
			t.Error(input, expected)
		} else {
			t.Log(input)
		}
	})
}

// === RUN   TestString
// === RUN   TestString/crack
// --- PASS: TestString (0.00s)
//     --- PASS: TestString/crack (0.00s)
//         fmtcraker_test.go:22: Str: hi
// PASS
// ok
