package problemSet

import (
	"fmt"
	"reflect"
	"testing"
)

func TestlongestCommonPrefix(t *testing.T) {
	tm := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"aaa", "aa", "aaa"}, "aa"},
		{[]string{"dog"}, "dog"},
		{[]string{"", "", ""}, ""},
		{[]string{"", ""}, ""},
	}
	for i, v := range tm {
		t.Run(fmt.Sprintf("Test%d\n", i), func(t *testing.T) {
			res := longestCommonPrefix(v.in)
			if !reflect.DeepEqual(res, v.want) {
				t.Errorf("want %q, but got %q\n", v.want, res)
			}
		})
	}
}
