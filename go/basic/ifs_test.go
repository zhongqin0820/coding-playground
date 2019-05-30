package basic

import (
	"testing"
)

func TestShow(t *testing.T) {
	rect := &Rect{width: 2, height: 3}
	cir := &Circle{radius: 3}
	Show("rect", rect)
	Show("cir", cir)
	Show("bool", true)
}
