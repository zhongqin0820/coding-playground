package basic

import (
	"log"
	"testing"
	"time"
)

// credit: https://www.zybuluo.com/phper/note/1072099
func server1(ch chan string) {
	time.Sleep(2e9)
	ch <- "server1"
}

func server2(ch chan string) {
	time.Sleep(1e9)
	ch <- "server2"
}

func TestSelect12(t *testing.T) {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		log.Println(s1)
	case s2 := <-output2:
		log.Println(s2)
	default:
		log.Println("no value")
	}
}
