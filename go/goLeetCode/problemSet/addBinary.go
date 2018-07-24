package problemSet

import (
	"log"
)

func AddBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		return b
	} else if lenB == 0 {
		return a
	}
	var tempA, tempB, res = []byte(a), []byte(b), []byte{}
	var carry bool
	var i, j int
	for i, j = lenA-1, lenB-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if tempA[i] == '0' && tempB[j] == '0' && carry == true {
			res = append(res, '1')
			carry = false
		} else if tempA[i] == '0' && tempB[j] == '0' && carry == false {
			res = append(res, '0')
		} else if tempA[i] == '1' && tempB[j] == '1' && carry == true {
			res = append(res, '1')
		} else if tempA[i] == '1' && tempB[j] == '1' && carry == false {
			res = append(res, '0')
			carry = true
		} else if tempA[i] != tempB[j] && carry == false {
			res = append(res, '1')
		} else if tempA[i] != tempB[j] && carry == true {
			res = append(res, '0')
			carry = false
		}
	}
	log.Printf("carry is : %t, i is : %d, j is : %d", carry, i, j)
	res = catTemp(res, tempA, i, &carry)
	res = catTemp(res, tempB, j, &carry)
	log.Printf("carry is : %t, res is %q\n", carry, res)
	if carry == true {
		res = append(res, '1')
	}
	res = reverse(res)
	log.Printf("res is : %q\n", res)
	return string(res)
}

func catTemp(s []byte, tempB []byte, j int, carry *bool) (res []byte) {
	log.Printf("s is :%q, tempB is %q, j is :%d, carry is :%t\n", s, tempB, j, *carry)
	for k := j; k >= 0; k-- {
		if tempB[k] == '0' && *carry == true {
			s = append(s, '1')
			*carry = false
		} else if tempB[k] == '0' && *carry == false {
			s = append(s, '0')
		} else if tempB[k] == '1' && *carry == true {
			s = append(s, '0')
		} else if tempB[k] == '1' && *carry == false {
			s = append(s, '1')
		}
	}
	log.Printf("s is %q\n", s)
	return s
}

func reverse(s []byte) (res []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// func main() {
// 	res := addBinary("110", "01")
// 	// res := addBinary("11", "1")
// 	// res := addBinary("1011", "1010")
// 	// res := addBinary("1", "111")
// 	// res := addBinary("100", "110010")
// 	// res := addBinary("1111", "1111")
// 	log.Println(res)
// }
