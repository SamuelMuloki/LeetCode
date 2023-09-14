package methods

import "fmt"

type Graph interface {
	AddEdge(v, w int)
	DFS(v int)
}

type DFSGraph struct {
	visited map[int]bool
	adj     map[int][]int
}

func NewDFSGraph() DFSGraph {
	return DFSGraph{
		visited: make(map[int]bool),
		adj:     make(map[int][]int),
	}
}

func (g DFSGraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
}

func (g DFSGraph) DFS(v int) {
	g.visited[v] = true
	fmt.Println(v)

	for _, val := range g.adj[v] {
		if !g.visited[val] {
			g.DFS(val)
		}
	}
}
