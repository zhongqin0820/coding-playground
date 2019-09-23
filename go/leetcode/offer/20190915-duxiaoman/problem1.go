package main

import (
	"fmt"
)

var diff1, diff2 int = -501, 501

type point struct {
	i, j int
}

// at indicates the value in point(i,j), false denotes illegal position
func (this point) at(maze [][]int) (int, bool) {
	if this.i < 0 || this.i >= len(maze) {
		return 0, true
	} else if this.j < 0 || this.j >= len(maze) {
		return 0, true
	}
	return maze[this.i][this.j], false
}

// directions
var points []point = []point{
	point{-1, 0}, //upward
	point{0, -1}, //leftward
	point{1, 0},  //downward
	point{0, 1},  //rightward
}

// next returns the next step point according to the directions
func (this point) next(neighbor point) point {
	neighbor.i += this.i
	neighbor.j += this.j
	return neighbor
}

// bread-first search the maze
func bfs(visited [][]int, maze [][]int, start, end point) [][]int {
	queue := []point{start}
	visited[start.i][start.j] = 1
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range points {
			n := node.next(neighbor)
			if v, ok := n.at(maze); ok || v == 1 {
				continue
			}
			if v, ok := n.at(visited); ok || v != 0 {
				continue
			}
			queue = append(queue, n)
			visited[n.i][n.j] = visited[node.i][node.j] + 1
		}
	}
	return visited
}

func problem1(x, y, n int, xs, ys []int) int {
	res := 0
	visited := make([][]int, diff2*2)
	maze := make([][]int, diff2*2)
	for i, _ := range visited {
		visited[i] = make([]int, diff2*2)
		maze[i] = make([]int, diff2*2)
	}
	for i := 0; i < n; i++ {
		maze[xs[i]+diff2][ys[i]+diff2] = 1
	}
	start := point{diff2, diff2}
	end := point{x + diff2, y + diff2}
	visited = bfs(visited, maze, start, end)
	res = visited[end.i][end.j] - 1
	return res
}

func problemWrapper1() {
	x, y, n, xs, ys := input1()
	output := problem1(x, y, n, xs, ys)
	fmt.Println(output)
}

func input1() (x, y, n int, xs, ys []int) {
	fmt.Scanf("%d %d %d", &x, &y, &n)
	xs, ys = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &xs[i], &ys[i])
	}
	return
}

func main() {
	problemWrapper1()
}

// 2 0 3
// 1 0
// 1 1
// 1 -1
