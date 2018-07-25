package main

import (
	// c "./contestSet"
	// p "./problemSet"
	"log"
)

func reverseString(s string) string {
	var b []byte = []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[j], b[i] = b[i], b[j]
	}
	return string(b)
}

func main() {
	// res := p.AddBinary("110010", "10111")
	// res := p.StrStr("hello", "ll")
	res := reverseString("hello")
	log.Printf("%d\n", res)
}
