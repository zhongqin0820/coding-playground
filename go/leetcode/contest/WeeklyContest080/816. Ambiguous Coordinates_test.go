package contest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-80/problems/ambiguous-coordinates/
func ambiguousCoordinates(S string) []string {
	res := make([]string, 0, len(S))
	s := S[1 : len(S)-1] // 去掉左右括号
	for i := 1; i < len(s); i++ {
		// 计算出以i为分割的左右两侧添加上小数点的合法数字
		lefts, rights := helper816AddDote(s[:i]), helper816AddDote(s[i:])
		for _, l := range lefts {
			for _, r := range rights {
				// 将左右进行拼接得到最终结果
				res = append(res, helper816Connect(l, r))
			}
		}
	}
	return res
}

// 判断这个切分是否是合法的浮点数
func helper816IsValid(s string) bool {
	f, _ := strconv.ParseFloat(s, 64)
	fs := strconv.FormatFloat(f, 'f', -1, 64)
	return s == fs
}

func helper816AddDote(s string) []string {
	res := make([]string, 0, len(s))
	if helper816IsValid(s) {
		res = append(res, s)
	}
	// 依次加小数点生成新的浮点数
	for i := 1; i < len(s); i++ {
		t := s[:i] + "." + s[i:]
		if helper816IsValid(t) {
			res = append(res, t)
		}
	}
	return res
}

// 连接合法的数值，形成坐标表示
func helper816Connect(left, right string) string {
	return fmt.Sprintf("(%s, %s)", left, right)
}

func TestAmbiguousCoordinates(t *testing.T) {
	tests := []struct {
		S      string
		output []string
	}{
		{"(123)", []string{"(1, 23)", "(1, 2.3)", "(12, 3)", "(1.2, 3)"}},
		{"(00011)", []string{"(0, 0.011)", "(0.001, 1)"}},
		{"(0123)", []string{"(0, 123)", "(0, 1.23)", "(0, 12.3)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"}},
		{"(100)", []string{"(10, 0)"}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, ambiguousCoordinates(ts.S))
		})
	}
}
