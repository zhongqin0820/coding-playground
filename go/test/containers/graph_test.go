package containers

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
)

type Node string

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

func (n *Node) String() string {
	return string(*n)
}

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

func (g *Graph) String() {
	g.lock.Lock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		neigh := g.edges[*g.nodes[i]]
		for j := 0; j < len(neigh); j++ {
			s += neigh[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.Unlock()
}

func (g *Graph) BFS(f func(*Node)) {
	g.lock.Lock()
	q := list.New()
	n := g.nodes[0]
	q.PushBack(n) //Element == *Node
	visited := make(map[*Node]bool)
	for q.Len() != 0 {
		// 出队列
		node := q.Front().Value.(*Node)
		q.Remove(q.Front())
		visited[node] = true
		neighs := g.edges[*node]
		// 将子节点加入队列
		for _, neigh := range neighs {
			if !visited[neigh] {
				q.PushBack(neigh)
				visited[neigh] = true
			}
		}
		// 打印
		if f != nil {
			f(node)
		}
	}
	g.lock.Unlock()
}

func TestGraph(t *testing.T) {
	var g Graph
	t.Run("Print", func(t *testing.T) {
		// A -> B C D
		// B -> A E
		// C -> A E
		// D -> A
		// E -> B C F
		// F -> E
		A, B, C, D, E, F := Node("A"), Node("B"), Node("C"), Node("D"), Node("E"), Node("F")
		g.AddNode(&A)
		g.AddNode(&B)
		g.AddNode(&C)
		g.AddNode(&D)
		g.AddNode(&E)
		g.AddNode(&F)
		g.AddEdge(&A, &B)
		g.AddEdge(&A, &C)
		g.AddEdge(&B, &E)
		g.AddEdge(&C, &E)
		g.AddEdge(&E, &F)
		g.AddEdge(&D, &A)
		g.String()
	})

	t.Run("BFS", func(t *testing.T) {
		// A
		// B
		// C
		// D
		// A
		// E
		// F
		g.BFS(func(n *Node) {
			fmt.Printf("%v\n", n.String())
		})
	})
}
