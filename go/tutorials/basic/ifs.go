package basic

import (
	"log"
	"math"
)

// Geometry is the interface of its area and perim
type Geometry interface {
	Area() float32
	Perim() float32
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perim() float32 {
	return 2 * math.Pi * c.radius
}

type Rect struct {
	width, height float32
}

func (r *Rect) Area() float32 {
	return r.height * r.width
}

func (r *Rect) Perim() float32 {
	return 2 * (r.height + r.width)
}

// Show all geometry info
func Show(name string, i interface{}) {
	switch i.(type) {
	case Geometry:
		log.Printf("%s's area=%.2f, perim=%.2f\n", name, i.(Geometry).Area(), i.(Geometry).Perim())
	default:
		log.Println("wrong type")
	}
}
