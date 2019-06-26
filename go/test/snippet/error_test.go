package snippet

import (
	"fmt"
	"reflect"
	"testing"
)

// 自定义结构体
type MyError1 struct {
	err string
}

// Error 返回错误说明
func (m MyError1) Error() string {
	errinfo := fmt.Sprintf("this is an Error no.3: %s ", m.err)
	return errinfo
}

// 重定义string
type MyError2 string

func (e MyError2) Error() string {
	return string(e)
}

// 比较
func TestMyError(t *testing.T) {
	// 自定义结构体
	m := MyError1{err: "My Error!"}
	t.Log(m)                         //this is an Error no.3: My Error!
	t.Log(m.Error())                 //this is an Error no.3: My Error!
	t.Log(reflect.TypeOf(m))         //*main.MyError1
	t.Log(reflect.TypeOf(m).Kind())  //struct
	t.Log(reflect.TypeOf(m.Error())) //string

	// 重定义
	var e MyError2 = "this is my error"
	t.Log(e)                         //this is my error
	t.Log(e.Error())                 //this is my error
	t.Log(reflect.TypeOf(e))         //main.MyError2
	t.Log(reflect.TypeOf(e).Kind())  //string
	t.Log(reflect.TypeOf(e.Error())) //string
}
