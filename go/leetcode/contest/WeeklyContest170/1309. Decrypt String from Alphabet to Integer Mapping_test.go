package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-170/problems/decrypt-string-from-alphabet-to-integer-mapping/
var helper1309 = map[string]string{
	"1":   "a",
	"2":   "b",
	"3":   "c",
	"4":   "d",
	"5":   "e",
	"6":   "f",
	"7":   "g",
	"8":   "h",
	"9":   "i",
	"10#": "j",
	"11#": "k",
	"12#": "l",
	"13#": "m",
	"14#": "n",
	"15#": "o",
	"16#": "p",
	"17#": "q",
	"18#": "r",
	"19#": "s",
	"20#": "t",
	"21#": "u",
	"22#": "v",
	"23#": "w",
	"24#": "x",
	"25#": "y",
	"26#": "z",
}

func freqAlphabets(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; {
		switch s[i] {
		case '#':
			res = helper1309[s[i-2:i+1]] + res
			i -= 3
		default:
			res = helper1309[s[i:i+1]] + res
			i -= 1
		}
	}
	return res
}

func TestFreqAlphabets(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"10#11#12", "jkab"},
		{"1326#", "acz"},
		{"25#", "y"},
		{"12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#", "abcdefghijklmnopqrstuvwxyz"},
		{"1", "a"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, freqAlphabets(ts.s))
		})
	}
}
