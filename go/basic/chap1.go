package basic

import (
	"fmt"
	"unicode/utf8"
)

type Q struct {
}

// Q2
func (q *Q) q2_1() {
	fmt.Println("q2_1")
	var count int = 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}
func (q *Q) q2_2() {
	fmt.Println("q2_2")
	var count int = 10
	var i int = 0
LABEL:
	if i < count {
		fmt.Println(i)
		i++
		goto LABEL
	}
}
func (q *Q) q2_3() {
	fmt.Println("q2_3")
	var count int = 10
	var i int = 0
	var a [10]int
	for ; i < count; i++ {
		a[i] = i
	}
	fmt.Println("%v", a)
}

// Q3
func (q *Q) q3() {
	fmt.Println("Fizz-Buzz")
	number := 0
LABEL:
	number++
	switch {
	case number > 100:
		return
	case number%3 == 0 && number%5 == 0:
		fmt.Println("FizzBuzz")
		goto LABEL
	case number%3 == 0:
		fmt.Println("Fizz")
		goto LABEL
	case number%5 == 0:
		fmt.Println("Buzz")
		goto LABEL
	default:
		fmt.Println(number)
		goto LABEL
	}
}

// Q4
func (q *Q) q4_1() {
	count := 100
	A := "A"
	for i := 0; i < count; i++ {
		fmt.Println(A)
		A += "A"
	}
}
func (q *Q) q4_2() {
	s := "asSASA ddd dsjkdsjs dk"
	sb := []byte(s)
	fmt.Println("%d,%d", len(sb), utf8.RuneCount(sb))
}
func (q *Q) q4_3() {
	// 只有将字符串转换成byte切片之后才可以交换
	s := "foobar"
	sb := []byte(s)
	i, j := 0, len(s)-1
	for ; i < j; i, j = i+1, j-1 {
		sb[i], sb[j] = sb[j], sb[i]
	}
	fmt.Println("%s", string(sb))
}

// Q5
func (q *Q) q5() {
	// 为了能够进行除法,必须将值转换为float64
	sf := []float64{1.1, 1.2, 1.7}
	var count float64
	for _, v := range sf {
		count += v
	}
	fmt.Println(count / float64(len(sf)))
}

// main function
func main1() {
	q := Q{}
	q.q5()
	// q.q4_3()
	// q.q4_2()
	// q.q4_1()
	// q.q3()
	// q.q2_1()
	// q.q2_2()
	// q.q2_3()
}
