package intro

import (
	"fmt"
	myexp "github.com/zhongqin0820/coding-playground/go/goIntro/myPack"
	"log"
	"math"
	"math/big"
	"regexp"
	"strconv"
)

// strconv regexp
func exp1() {
	s := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(s)); ok {
		log.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(s, "##.#")
	log.Println(str)
	str = re.ReplaceAllStringFunc(s, f)
	log.Println(str)
	return
}

// math math/big
func exp2() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	log.Println(ip)
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	log.Println(rq)
	return
}

func exp3() {
	log.Println(myexp.MyVar)
	myexp.MyFunc()
}

func main9() {
	// exp1()
	// exp2()
	exp3()
	fmt.Println("...")
}
