package main

import (
	"fmt"
	"strings"
)

func exp1() {
	// var a int
	var b int32
	// b = a + a // bugs here, can't use (type int) as type int32
	b = b + 5
	fmt.Println(b)
	return
}

func exp2() {
	fmt.Println(`hahahahah\n\`)
}

// count_characters.go count a string's byte and rune size
func exp3() {
	// need more implementation here
	s := "aaSASA ddd dsjkdsjs dk"
	fmt.Println(len(s))
}

func exp4() {
	s := "test about strings package suffix"
	res := strings.HasPrefix(s, "test")
	fmt.Println("HasPrefix: ", res)
	res = strings.HasSuffix(s, "suffix")
	fmt.Println("HasSuffix: ", res)
	return
}

func main() {
	// exp1()
	// exp2()
	// exp3()
	exp4()
	fmt.Println("......")
}
