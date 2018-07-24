package problemSet

import (
	"log"
	"testing"
)

// res := addBinary("110", "01")
// res := addBinary("11", "1")
// res := addBinary("1011", "1010")
// res := addBinary("1", "111")
// res := addBinary("100", "110010")
// res := addBinary("1111", "1111")
func TestAddBinary(t *testing.T) {
	if res := AddBinary("110", "01"); res != "111" { //try a unit test on function
		log.Printf("%q\n", res)
		t.Error("110 + 01 没有通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}
