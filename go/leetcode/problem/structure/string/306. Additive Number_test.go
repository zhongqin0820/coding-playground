package problem

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/additive-number/
func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	n := len(num)
	for i := 1; i < n; i++ {
		if i > 1 && num[0] == '0' {
			break
		}
		for j := i + 1; j < n; j++ {
			if j-i > 1 && num[i] == '0' { // 去掉num2出现leading zeros的情况
				break
			}
			s := num[0:j]
			num1, _ := strconv.Atoi(num[0:i])
			num2, _ := strconv.Atoi(num[i:j])
			// 往前迭代s
			for len(s) <= n {
				temp := strconv.Itoa(num1 + num2)
				s += temp
				num1 = num2
				num2, _ = strconv.Atoi(temp)
				// fmt.Println(s)
				if strings.Compare(s, num) == 0 {
					return true
				}
			}
		}
	}
	return false
}

func TestIsAdditiveNumber(t *testing.T) {
	tests := []struct {
		num    string
		output bool
	}{
		{"112358", true},
		{"199100199", true},
		{"1023", false},
		{"101", true},
		{"000", true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isAdditiveNumber(ts.num))
		})
	}
}
