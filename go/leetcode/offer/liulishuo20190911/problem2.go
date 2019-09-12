package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type point struct {
	i, j int
}

// at indicates the value in point(i,j), false denotes illegal position
func (this point) at(maze [][]int) (int, bool) {

	if this.i < 0 || this.i >= len(maze) {
		return 0, true
	} else if this.j < 0 || this.j >= len(maze[0]) {
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

func problem2(maze [][]int, start, end point) int {
	res := 0
	visited := make([][]int, len(maze))
	for i, _ := range visited {
		visited[i] = make([]int, len(maze[0]))
	}
	visited = bfs(visited, maze, start, end)
	res = visited[end.i][end.j] - 1
	return res
}

func problemWrapper2() {
	A, start, end := input2()
	// fmt.Println(A, start, end)
	output := problem2(A, start, end)
	fmt.Println(output)
}

func input2() ([][]int, point, point) {
	var input io.Reader
	input = os.Stdin
	reader := bufio.NewReader(input)
	scanner := bufio.NewScanner(reader)
	MAXLEN := 1024
	A := make([][]int, 0, MAXLEN)
	var i int
	var start, end point
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 || s == "\n" {
			break
		}
		tempS := strings.Split(strings.TrimRight(s, " "), " ")
		if len(tempS) == 0 {
			break
		}
		temp := []int{}
		for j, v := range tempS {
			num, _ := strconv.Atoi(v)
			temp = append(temp, num)
			if num == 2 {
				start = point{i, j}
			}
			if num == 3 {
				end = point{i, j}
			}
		}
		i++
		A = append(A, temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
	return A, start, end
}

// func main() {
// 	problemWrapper2()
// }
