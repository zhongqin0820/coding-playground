package goTest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tm := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 9}, 10},
	}
	for i, v := range tm {
		t.Run(fmt.Sprintf("Test%d", i), func(t *testing.T) {
			res := Sum(v.in)
			if !reflect.DeepEqual(res, v.want) {
				t.Errorf("want %d, got %d\n", v.want, res)
			}
		})
	}
}

func TestSumAdd(t *testing.T) {
	tm := []struct {
		in1, in2, want []int
	}{
		{[]int{}, []int{}, []int{0, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 9}, []int{15, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{}, []int{15, 0}},
		{[]int{1, 2}, []int{-2, 9}, []int{3, 7}},
	}
	for i, v := range tm {
		t.Run(fmt.Sprintf("Test %d\n", i), func(t *testing.T) {
			res := SumAdd(v.in1, v.in2)
			if !reflect.DeepEqual(res, v.want) {
				t.Errorf("want %v, got %v\n", v.want, res)
			}
		})
	}
}

func TestSumTails(t *testing.T) {
	tm := []struct {
		in1, in2, want []int
	}{
		{[]int{}, []int{}, []int{0, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 9}, []int{14, 9}},
		{[]int{1, 2, 3, 4, 5}, []int{}, []int{14, 0}},
		{[]int{1, 2}, []int{-2, 9}, []int{2, 9}},
	}
	for i, v := range tm {
		t.Run(fmt.Sprintf("Test %d\n", i), func(t *testing.T) {
			res := SumTails(v.in1, v.in2)
			if !reflect.DeepEqual(res, v.want) {
				t.Errorf("want %v, got %v\n", v.want, res)
			}
		})
	}
}
