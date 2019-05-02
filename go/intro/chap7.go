package intro

import (
	"bytes"
	"fmt"
	myexp "github.com/zhongqin0820/coding-playground/go/goIntro/myPack"
	"log"
	"strings"
)

func exp1() {
	var buffer bytes.Buffer
	s := strings.Fields("test1 test2 test3")
	for _, v := range s {
		buffer.WriteString(v)
	}
	log.Printf("%v", buffer.String())
}

func exp2() {
	s1 := []byte{'p', 'o', 'e', 'm'}
	log.Printf("%q", s1)
	s2 := s1[2:]
	s2[1] = 't'
	log.Printf("%q %q", s1, s2)
}

func exp3() {
	myexp.Hi()
	log.Println(myexp.MyPackA)
}

func main7() {
	// exp1()
	// exp2()
	exp3()
	fmt.Println("...")
}
