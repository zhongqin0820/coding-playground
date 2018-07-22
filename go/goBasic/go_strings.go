package main

import (
	"fmt"
	"strconv"
	"strings"
)

// string manipulation
func StringManipulation() {
	s := []string{"fkalg", "sdalgkjh", "akjgbaku"}
	// 字符串拼接
	ss := strings.Join(s, " ")
	fmt.Println(ss)
	// 查找
	fmt.Println(strings.Contains(ss, s[0]))
	// 查找
	if pos := strings.Index(ss, s[0]); pos == -1 {
		fmt.Println("cann't find")
	} else {
		fmt.Println("pos at :", pos)
	}
	// 重复
	sss := strings.Repeat(s[0], 2)
	fmt.Println(sss)
	// 替换
	a := strings.Replace(ss, s[0], "new", 3)
	fmt.Println(a)
	// 分割
	b := strings.Split(ss, " ")
	fmt.Printf("%q\n", b)
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	// 去掉头尾
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! "))
	// Fields
	fmt.Printf("Fields are: %q\n", strings.Fields(ss))
}

// strconv字符转换
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	// 注意str 是字节数组
	// for _, v := range str {
	// 	fmt.Printf("%q ", v)
	// }
	fmt.Println(string(str))
	//Format 系列函数把其他类型的转换为字符串
	a1 := strconv.FormatBool(false)
	b1 := strconv.FormatFloat(123.23, 'g', 12, 64)
	c1 := strconv.FormatInt(1234, 10)
	d1 := strconv.FormatUint(12345, 10)
	e1 := strconv.Itoa(1023)
	fmt.Println(a1, b1, c1, d1, e1)
	fmt.Printf("%T %T %T %T %T\n", a1, b1, c1, d1, e1)
	//Parse 系列函数把字符串转换为其他类型
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
	fmt.Printf("%T %T %T %T %T\n", a, b, c, d, e)
}
