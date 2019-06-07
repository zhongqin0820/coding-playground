package adapter

import (
	"log"
	"testing"
)

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
	log.Print("this is a testing")
}
