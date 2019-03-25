package problemSet

import (
	"log"
)

func StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}

func strStr(haystack string, needle string) int {
	h := []byte(haystack)
	n := []byte(needle)
	// needle is empty
	if len(n) == 0 {
		return 0
	}
	// haystack is shorter than needle
	lenN, lenH := len(n), len(h)
	if lenN > lenH {
		return -1
	}
	for i, _ := range h {
		if lenH-i < lenN {
			break
		}
		log.Printf("%s %s\n", string(h[i:]), string(n))
		if meetEqual(h[i:], n) == true {
			return i
		}
	}
	return -1
}

func meetEqual(b, n []byte) bool {
	if len(b) == 0 || len(n) == 0 {
		return false
	}
	lenB := len(b)
	for i, v := range n {
		if i < lenB && v == b[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
