package problem

import (
	"log"
)

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	lenS := len(strs)
	if lenS == 0 {
		return ""
	} else if lenS == 1 {
		return strs[0]
	}
	var res []byte = []byte(strs[0])
	for i := 1; i < lenS; i++ {
		a := []byte(strs[i])
		if res = meetUnequal(a, res); len(res) == 0 {
			return ""
		}
	}
	return string(res)
}

// return the prefix of two slice of byte
func meetUnequal(a, b []byte) []byte {
	lenA, lenB, ind := len(a), len(b), 0
	if lenA == 0 || lenB == 0 {
		return []byte("")
	}
	for i, v := range a {
		if i >= lenB {
			break
		}
		if v != b[i] && i == 0 {
			return []byte("")
		} else if v != b[i] {
			return b[0:i]
		}
		ind = i
	}
	log.Printf("%s %s %d\n", string(a), string(b), ind)
	return a[:ind+1]
}
