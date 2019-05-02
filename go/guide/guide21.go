package guide

import (
	"fmt"
	"math"
	"os"
	_ "strconv"
	"time"
)

// method is defined in struct
type Vertex struct {
	X float64
	Y float64
}

// interface is defined in interface
type Abser interface {
	Abs() float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

type MyFloat float64

func (m MyFloat) Abs() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

func exp1() {
	v := &Vertex{2, 2}
	fmt.Printf("%f\n", v.Abs())
}

func exp2() {
	v := &Vertex{2, 2}
	fmt.Printf("%f\n", v.Abs())
	v.Scale(4)
	fmt.Printf("%f\n", v.Abs())
}

func exp3() {
	v := &Vertex{1, 1}
	fmt.Printf("%f\n", v.Abs())
	var a Abser
	a = v
	fmt.Printf("%f\n", a.Abs())
	m := MyFloat(4)
	fmt.Println(m.Abs())
	a = m
	fmt.Println(a.Abs())
	fmt.Printf("%T", a)
	a = &m
	// error here, MyFloat implements Abs() not through pointer
	fmt.Println(a.Abs())
	fmt.Printf("%T", a)
}

type Reader interface {
	Read(r []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func exp4() {
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func exp5() {
	a := Person{"A", 12}
	b := Person{"B", 21}
	fmt.Println(a, b)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	// var s string
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func exp6() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s ", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work!",
	}
}

func exp7() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func exp8() {
	r, e := Sqrt(8.0)
	if e == nil {
		fmt.Println(r)
	}
	r, e = Sqrt(-8.0)
	if e == nil {
		fmt.Println(r)
	}
}

func main2() {
	// exp1()
	// exp2()
	// exp3()
	// exp4()
	// exp5()
	// exp6()
	// exp7()
	exp8()
}
