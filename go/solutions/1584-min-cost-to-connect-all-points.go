package solutions

import "container/heap"

type Edge struct {
	Vertex1, Vertex2, Weight int
}

func MinCostConnectPoints(points [][]int) int {
	parent := make([]int, len(points))
	rank := make([]int, len(points))

	for i := 0; i < len(points); i++ {
		parent[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if i != parent[i] {
			parent[i] = find(parent[i])
		}

		return parent[i]
	}

	union := func(v1, v2 int) bool {
		p1, p2 := find(v1), find(v2)

		if p1 == p2 {
			return false
		}

		switch {
		case rank[p1] < rank[p2]:
			parent[p1] = p2
		case rank[p1] > rank[p2]:
			parent[p2] = p1
		default:
			parent[p2] = p1
			rank[p1]++

		}

		return true
	}

	abs := func(val int) int {
		if val < 0 {
			return -val
		}

		return val
	}

	getWeight := func(pointX, pointY []int) int {
		return abs(pointX[0]-pointY[0]) + abs(pointX[1]-pointY[1])
	}

	edges := &EdgeMinHeap{}
	heap.Init(edges)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			heap.Push(edges, Edge{
				Vertex1: i,
				Vertex2: j,
				Weight:  getWeight(points[i], points[j]),
			})
		}
	}

	var output int
	connected := 0

	for edges.Len() > 0 {
		edge := heap.Pop(edges).(Edge)
		if union(edge.Vertex1, edge.Vertex2) {
			connected++
			output += edge.Weight
		}

		if connected == len(points)-1 {
			return output
		}
	}

	return output
}

type EdgeMinHeap []Edge

func (h EdgeMinHeap) Len() int           { return len(h) }
func (h EdgeMinHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h EdgeMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *EdgeMinHeap) Push(x any)        { *h = append(*h, x.(Edge)) }
func (h *EdgeMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
