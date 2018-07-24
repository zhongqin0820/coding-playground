package main

import (
	// c "./contestSet"
	p "./problemSet"
	"log"
)

func main() {
	res := p.AddBinary("1", "11")
	log.Printf("%q\n", res)
}
