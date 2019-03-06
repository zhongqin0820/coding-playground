package singleton

import (
	"fmt"
	"testing"
)

func Test_GetInstance(t *testing.T) {
	s := GetInstance()
	s.AddOne()
	fmt.Println(s.GetValue())
	fmt.Println(GetInstance().AddOne())
}
