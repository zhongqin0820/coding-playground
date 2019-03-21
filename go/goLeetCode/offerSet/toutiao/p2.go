package xmain

import (
	"fmt"
)

func DelRep(b []byte) []byte {
	if len(b) <= 2 {
		return b
	}
	if len(b) == 3 && b[0] == b[1] && b[1] == b[2] {
		return b[0:2]
	}
	var i, j, k, l = 0, 1, 2, 3

	for len(b) > 2 {
		// AAA
		if b[i] == b[j] && b[j] == b[k] {
			b = append(b[0:k], b[k+1:]...)
		} else if len(b) > 3 && l < len(b) {
			//AABB
			if b[i] == b[j] && b[k] == b[l] {
				b = append(b[0:k], b[k+1:]...)
			} else {
				i, j, k, l = j, k, l, l+1
			}
		} else {
			break
		}
	}
	return b
}

func xmain() {
	N := 0
	fmt.Scanf("%d", &N)
	for ; N > 0; N-- {
		s := ""
		fmt.Scan(&s)
		if len(s) == 0 {
			continue
		}
		b := []byte(s)
		c := DelRep(b)
		fmt.Println(string(c))
	}
}
