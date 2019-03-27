package problem

import (
	"testing"
)

// res := addBinary("110", "01")
// res := addBinary("11", "1")
// res := addBinary("1011", "1010")
// res := addBinary("1", "111")
// res := addBinary("100", "110010")
// res := addBinary("1111", "1111")
// "110010"
// "10111"
func TestaddBinary(t *testing.T) {
	var i int = 1
	if res := addBinary("110", "01"); res != "111" { //try a unit test on function
		t.Logf("%q\n", res)
		t.Errorf("第%d个测试没有通过", i) // 如果不是如预期的那么就报错
	} else {
		t.Logf("第%d个测试通过了", i) //记录一些你期望记录的信息
	}
	i++
	if res := addBinary("11", "1"); res != "100" {
		t.Logf("%q\n", res)
		t.Errorf("第%d个测试没有通过", i)
	} else {
		t.Logf("第%d个测试通过了", i)
	}
	i++
	if res := addBinary("1011", "1010"); res != "10101" {
		t.Logf("%q\n", res)
		t.Errorf("第%d个测试没有通过", i)
	} else {
		t.Logf("第%d个测试通过了", i)
	}
	i++
	if res := addBinary("1", "111"); res != "1000" {
		t.Logf("%q\n", res)
		t.Errorf("第%d个测试没有通过", i)
	} else {
		t.Logf("第%d个测试通过了", i)
	}
	i++
	if res := addBinary("1111", "1111"); res != "11110" {
		t.Logf("%q\n", res)
		t.Errorf("第%d个测试没有通过", i)
	} else {
		t.Logf("第%d个测试通过了", i)
	}
	i++
	if res := addBinary("110010", "10111"); res != "1001001" {
		t.Logf("%q\n", res)
		t.Errorf("第%d个测试没有通过", i)
	} else {
		t.Logf("第%d个测试通过了", i)
	}
	i++
}
