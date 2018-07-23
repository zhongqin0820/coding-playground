package main

import (
	"fmt"
	"log"
)

func exp1() {
	log.Println("start")
	defer func() {
		if e := recover(); e != nil {
			log.Println("fix?", e)
		}
	}()
	panic("error")
	log.Println("end")
}

func main() {
	exp1()
	fmt.Println("......")
}
