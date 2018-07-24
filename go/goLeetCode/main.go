package main

import (
	// c "./contestSet"
	p "./problemSet"
	"log"
)

func main() {
	res := p.AddBinary("110010", "10111")
	log.Printf("%q\n", res)
}
