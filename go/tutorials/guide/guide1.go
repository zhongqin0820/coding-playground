package guide

import (
	"fmt"
	// "golang.org/x/tour/pic"
	_ "math"
	"runtime"
	"strings"
	"time"
)

// package,variable,function
func MySqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

// for,if,else,switch
func PrintOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

func WhenSaturday() {
	switch today := time.Now().Weekday(); time.Saturday {
	case today:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

func GreetingTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func TestDefer() {
	defer fmt.Println("World")
	fmt.Println("Hello1")
	fmt.Println("Hello2")
	fmt.Println("Hello3")
}

func TestDeferStack() {
	fmt.Println("start")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("defer %d", i)
	}
	fmt.Println("over")
}

// pointer,struct,slice,map
func UsePointer() {
	i, _ := 1, 2
	p := &i
	fmt.Println(*p)
	*p = 3
	fmt.Println(i)
}

func UseStruct() {
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	fmt.Println(v.X, v.Y)
}

func UseArray() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func UseSlice() {
	// p := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println("p=", p)
	// lenp := len(p)
	// for i := 0; i < lenp; i++ {
	// 	fmt.Printf("p[%d]=%d\n", i, p[i])
	// }
	// a := make([]int, 5)
	var a []int
	fmt.Printf("%d %d %v\n", len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("%d %d %v\n", len(a), cap(a), a)
}

func UseRange() {
	var a []int = []int{1, 2, 4, 8, 16}
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	return
}

// func Pic(dx, dy int) [][]uint8 {
// 	a := make([][]uint8, dy)
// 	for x := range a {
// 		b := make([]uint8, dx)
// 		for y := range b {
// 			b[y] = uint8(x*y - 1)
// 		}
// 		a[x] = b
// 	}
// 	return a
// }

func UseMap() {
	type Vertex struct {
		X int
		Y int
	}
	// declaration
	var m map[string]Vertex
	// assign memory
	m = make(map[string]Vertex)
	m["P1"] = Vertex{1, 2}
	fmt.Println(m["P1"])
	delete(m, "P1")
	elem, ok := m["P1"]
	fmt.Println(ok)
	if ok == true {
		fmt.Println(elem)
	} else {
		m["P1"] = Vertex{2, 3}
	}
	fmt.Println(m["P1"])
}

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	a := strings.Fields(s)
	fmt.Println(a)
	for _, v := range a {
		res[v]++
	}
	return res
}

func Fibonacci() func() int {
	i := 0
	j := 1
	return func() int {
		temp := i
		i, j = j, i+j
		return temp
	}
}

// main function
func main1() {
	// fmt.Println(math.Sqrt(5))
	// fmt.Println(MySqrt(5))
	// PrintOS()
	// WhenSaturday()
	// GreetingTime()
	// defer TestDefer()
	// fmt.Println("testing defer!")
	// TestDeferStack()
	// UsePointer()
	// UseStruct()
	// UseArray()
	// UseSlice()
	// UseRange()
	// pic.Show(Pic)
	// UseMap()
	// fmt.Println(WordCount("haha me"))
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
