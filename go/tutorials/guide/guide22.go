package guide

import (
	"fmt"
	// "golang.org/x/tour/reader"
	// "golang.org/x/tour/pic"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// slice init
func exp1() {
	var e [][]float64
	n := 4
	e = make([][]float64, n)
	fmt.Printf("%T\n%v", e, e)
	/*
		[][]float64
		[[] [] [] []]
	*/
}

func exp2() {
	r := strings.NewReader("Hello Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// type MyReader struct{}

// func (m MyReader) Read(b []byte) (int, error) {
// 	b[0] = 'A'
// 	return 1, nil
// }

// func exp3() {
// 	reader.Validate(MyReader{})
// }

type rot13Reader struct {
	r io.Reader
}

func rot13(p byte) byte {
	switch {
	case p >= 'A' && p <= 'M':
		p += 13
	case p >= 'N' && p <= 'Z':
		p -= 13
	case p >= 'a' && p <= 'm':
		p += 13
	case p >= 'n' && p <= 'z':
		p -= 13
	}
	return p
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	ori := make([]byte, 50)
	i, err := r.r.Read(ori)
	for index, value := range ori[:i] {
		p[index] = rot13(value)
	}
	return i, err
}

func exp4() {
	s := strings.NewReader("Lbh penpxafhka akjbf!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!")
}

func exp5() {
	var h Hello
	err := http.ListenAndServe("localhost:9999", h)
	if err != nil {
		log.Fatal(err)
	}
}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res string
	res = s.Greeting + s.Punct + s.Who
	fmt.Fprint(w, res)
}

func exp6() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// package image
func exp7() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

type Image struct{}

func (i *Image) colorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}

func exp8() {
	m := Image{}
	pic.ShowImage(m)
}

func main22() {
	// exp1()
	// exp2()
	// exp3()
	// exp4()
	// exp5()
	// exp6()
	exp7()
}
