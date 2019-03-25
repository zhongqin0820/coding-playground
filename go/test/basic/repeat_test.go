package basic

import (
	"fmt"
	"testing"
)

func TestRepeats(t *testing.T) {
	ts := []struct {
		ch   string // the character
		ind  int    // repeat times
		want string // result wanted
	}{
		{"a", 10, "aaaaaaaaaa"},
		{"aa", 5, "aaaaaaaaaa"},
	}
	for i, v := range ts {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			res := Repeat(v.ch, v.ind)
			if res != v.want {
				t.Errorf("want %s, but got %s\n", v.want, res)
			}
		})
	}
}

// go test -bench=.
func BenchmarkRepeats(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a", i)
	}
}

// Repeat a string n times
func ExampleRepeats() {
	res := Repeat("a", 5)
	fmt.Println(res)
	// Output: aaaaa
}
