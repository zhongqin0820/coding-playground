package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func angleClock(hour int, minutes int) float64 {
	// 分钟形成固定角度: minutes / 60 * 360
	// 时钟下限形成固定角度，时钟在间隔之间还有一个角度：hour / 12 * 360 + minutes / 60 * 30
	// 15 / 60 * 360 = 90
	// 3 / 12 * 360 + 15 / 60 * 30 = 97.5
	var res float64
	m := float64(minutes * 6)
	h := float64((hour%12)*30) + float64(minutes)/2.0
	switch {
	case m <= 180 && h <= 180:
		res = max(m-h, h-m)
	case m >= 180 && h >= 180:
		res = max(m-h, h-m)
	case m >= 180 && h <= 180:
		res = min(m-h, min(m, 360-m)+h)
	case m <= 180 && h >= 180:
		res = min(h-m, min(h, 360-h)+m)
	}
	if res > 180 {
		return res - 180
	}
	return res
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func TestAngleClock(t *testing.T) {
	tests := []struct {
		hour    int
		minutes int
		output  float64
	}{
		{1, 57, 76.5},
		{12, 30, 165},
		{3, 30, 75},
		{3, 15, 7.5},
		{4, 50, 155},
		{12, 0, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, angleClock(ts.hour, ts.minutes))
		})
	}
}
