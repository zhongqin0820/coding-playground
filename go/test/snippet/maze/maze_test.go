package maze

import (
	"fmt"
	"os"
	"testing"
)

// LoadData from the file, first row indicates the row & the col of the maze
// the rest of the data is the maze
func LoadData(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	row, col := 0, 0
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i, _ := range maze {
		maze[i] = make([]int, col)
	}
	for i, _ := range maze {
		for j, _ := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

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
		//
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

func TestMaze(t *testing.T) {
	const fileName = "maze.in"
	maze := LoadData(fileName)
	visited := make([][]int, len(maze))
	for i, _ := range visited {
		visited[i] = make([]int, len(maze[0]))
	}
	visited = bfs(visited, maze, point{0, 0}, point{len(maze), len(maze[0])})
	for i, _ := range visited {
		for _, v := range visited[i] {
			fmt.Printf("%3d", v)
		}
		println()
	}
}
