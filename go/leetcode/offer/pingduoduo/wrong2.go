package main

import (
	"fmt"
	"strings"
)

func judgeN(n string) bool {
	if len(n) == 0 || len(n) > 52 {
		return false
	}
	for _, v := range []byte(n) {
		if (v-'A' >= 0 && v-'A' <= 26) || (v-'a' >= 0 && v-'a' <= 26) {
			return true
		}
		return false
	}
	return true
}

func toNum(v byte) int {
	if v-'A' >= 0 && v-'A' <= 26 {
		return int(v - 'A')
	}
	return int(v - 'a')
}

func main() {
	var n string
	fmt.Scan(&n)
	if judgeN(n) == false {
		return
	}
	n = strings.ToLower(n)
	var flag = 0
	var i, j = 1, len(n) - 1
	for i <= j {
		if toNum(n[i]) < toNum(n[flag]) {
			var change bool
			for k := j; k > i; k-- {
				if toNum(n[k]) == toNum(n[flag]) {
					flag = i
					i++
					change = true
					break
				}
			}
			if change == false {
				break
			}
		} else {
			i++
		}
	}
	fmt.Printf("%c\n", n[flag])
}
