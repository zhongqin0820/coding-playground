package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/walking-robot-simulation/
func robotSim(commands []int, obstacles [][]int) int {
	if commands == nil || len(commands) == 0 {
		return 0
	}
	size := len(obstacles)
	m := make(map[string]bool, size)
	for _, obs := range obstacles {
		m[fmt.Sprintf("%d,%d", obs[0], obs[1])] = true
	}
	x, y, res := 0, 0, 0
	i, size_n := 0, len(commands)

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	j := 0
	// 0, 1  : north
	// 1, 0  : right
	// 0, -1 : south
	// -1, 0 : left

	for ; i < size_n; i++ {
		switch commands[i] {
		case -2: //turn left 90 degrees
			j = (j + 3) % 4
		case -1: //turn right 90 degrees
			j = (j + 1) % 4
		default:
			//1 <= x <= 9: move forward x units
			for step := 0; step < commands[i] && !m[fmt.Sprintf("%d,%d", x+dirs[j][0], y+dirs[j][1])]; step++ {
				x += dirs[j][0]
				y += dirs[j][1]
			}
		}
		if res < x*x+y*y {
			res = x*x + y*y
		}
	}
	return res
}

func TestRobotSim(t *testing.T) {
	tests := []struct {
		commands  []int
		obstacles [][]int
		output    int
	}{
		{[]int{4, -1, 3}, [][]int{}, 25},
		{[]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}, 65},
		{[]int{-2, 8, 3, 7, -1}, [][]int{{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3}}, 324},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, robotSim(ts.commands, ts.obstacles))
		})
	}
}
