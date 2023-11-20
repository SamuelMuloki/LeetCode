package solutions

import "container/heap"

type Graph struct {
	gr map[int][][2]int
	n  int
}

func GraphConstructor(n int, edges [][]int) Graph {
	gr := make(map[int][][2]int, n)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		gr[u] = append(gr[u], [2]int{w, v})
	}
	return Graph{
		gr: gr,
		n:  n,
	}
}

func (g *Graph) AddEdge(edge []int) {
	u, v, w := edge[0], edge[1], edge[2]
	g.gr[u] = append(g.gr[u], [2]int{w, v})
}

const inf = 1 << 31

func (g *Graph) ShortestPath(node1 int, node2 int) int {
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = inf
	}
	dist[node1] = 0
	h := &PathMinHeap{}
	heap.Init(h)
	h.Push([2]int{0, node1})

	used := make([]bool, g.n)

	for h.Len() > 0 {
		edge := heap.Pop(h).([2]int)
		if used[edge[1]] {
			continue
		}
		if edge[1] == node2 {
			return dist[node2]
		}
		used[edge[1]] = true

		conns := g.gr[edge[1]]
		for i := range conns {
			ne := conns[i]
			if dist[edge[1]]+ne[0] < dist[ne[1]] {
				dist[ne[1]] = dist[edge[1]] + ne[0]
				heap.Push(h, [2]int{dist[ne[1]], ne[1]})
			}
		}
	}
	return -1
}

type PathMinHeap [][2]int

func (h PathMinHeap) Len() int            { return len(h) }
func (h PathMinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h PathMinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PathMinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *PathMinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}
