package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func problem(a, b, c, d, x, y, x1, y1, x2, y2 int) string {
	sp := b - a
	cz := d - c
	if sp >= 0 && x2-x < sp {
		return "No"
	} else if x-x1 < (-sp) {
		return "No"
	}
	if cz >= 0 && y2-y < cz {
		return "No"
	} else if y-y1 < (-cz) {
		return "No"
	}
	if ((a > 0 || b > 0) && x2-x1 == 0) || ((c > 0 || d > 0) && y2-y1 == 0) {
		return "No"
	}
	return "Yes"
}

// data preprocessing
func input() (a, b, c, d, x, y, x1, y1, x2, y2 int) {
	fmt.Scanf("%d %d %d %d\n", &a, &b, &c, &d)
	fmt.Scanf("%d %d %d %d %d %d\n", &x, &y, &x1, &y1, &x2, &y2)
	return
}

func problemWrapper() {
	t := 0
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		a, b, c, d, x, y, x1, y1, x2, y2 := input()
		res := problem(a, b, c, d, x, y, x1, y1, x2, y2)
		fmt.Println(res)
	}
}

func main() {
	problemWrapper()
}

// local testing
func TestProblem(t *testing.T) {
	tests := []struct {
		// 𝑎  to the left, 𝑏 to the right, 𝑐 to the down, 𝑑 to the up
		a, b, c, d int
		// 𝑥1≤𝑢≤𝑥2, 𝑦1≤𝑣≤𝑦2.
		x, y, x1, y1, x2, y2 int
		output               string
	}{
		{3, 2, 2, 2, 0, 0, -2, -2, 2, 2, "Yes"},
		{3, 1, 4, 1, 0, 0, -1, -1, 1, 1, "No"},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, "No"},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, "Yes"},
		{5, 1, 1, 1, 0, 0, -100, -100, 0, 100, "Yes"},
		{1, 1, 5, 1, 0, 0, -100, -100, 100, 0, "Yes"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem(ts.a, ts.b, ts.c, ts.d, ts.x, ts.y, ts.x1, ts.y1, ts.x2, ts.y2))
		})
	}
}
