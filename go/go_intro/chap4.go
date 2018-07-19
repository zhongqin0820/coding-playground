package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
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
	s := "test about strings package test suffix"
	res := strings.HasPrefix(s, "test")
	fmt.Println("HasPrefix: ", res)
	res = strings.HasSuffix(s, "suffix")
	fmt.Println("HasSuffix: ", res)
	res = strings.Contains(s, "test")
	log.Println("Contains", res)
	resInt := strings.Index(s, "test")
	log.Println("Index", resInt)
	// log.Fatalln("Index", resInt)
	// log.Panicln("Index", resInt)
	resInt = strings.LastIndex(s, "test")
	log.Println("Index", resInt)
	resString := strings.Replace(s, "es", "hhhhhhhhh", 1)
	log.Println("Replace", resString)
	resInt = strings.Count(s, "test")
	log.Println("Count", resInt)
	resString = strings.Repeat("test ", 3)
	log.Println(resString)
	resString = strings.ToUpper(resString)
	log.Println(resString)
	resString = strings.ToLower(resString)
	log.Println(resString)
	testString := "     resr   rserrrrrrrrrrrrrrr     "
	testString = strings.TrimSpace(testString)
	log.Println(testString)
	testString = strings.Trim(testString, "r")
	log.Println(testString)
	resSlice := strings.Fields(s)
	log.Printf("%v", resSlice)
	log.Println(len(resSlice))
	resSlice = strings.Split(s, " ")
	log.Printf("%v", resSlice)
	resString = strings.Join(resSlice, ",")
	log.Println(resString)
	convInt, err := strconv.Atoi("123")
	if err != nil {
		log.Printf("%T", err)
	} else {
		log.Println(convInt, err)
	}
	t := time.Now()
	log.Println(t, "\n", t.Day(), t.Month(), t.Hour())
	return
}

func main() {
	// exp1()
	// exp2()
	// exp3()
	exp4()
	fmt.Println("......")
}
