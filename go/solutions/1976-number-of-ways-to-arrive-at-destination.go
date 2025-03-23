package solutions

import (
	"container/heap"
	"math"
)

type countPathsEdge struct {
	dst  int
	time int64
}

type Item struct {
	node int
	time int64
	ways int64
}

func CountPaths(n int, roads [][]int) int {
	graph := make([][]countPathsEdge, n)
	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], countPathsEdge{road[1], int64(road[2])})
		graph[road[1]] = append(graph[road[1]], countPathsEdge{road[0], int64(road[2])})
	}

	dist, ways := make([]int64, n), make([]int64, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0
	ways[0] = 1

	h := &CountPathsHeap{}
	heap.Init(h)
	heap.Push(h, Item{0, 0, 1})

	for h.Len() > 0 {
		curr := heap.Pop(h).(Item)
		u := curr.node

		if curr.time > dist[u] {
			continue
		}

		for _, edge := range graph[u] {
			v := edge.dst
			newTime := dist[u] + edge.time
			if newTime < dist[v] {
				dist[v] = newTime
				ways[v] = ways[u]
				heap.Push(h, Item{v, newTime, ways[v]})
			} else if newTime == dist[v] {
				ways[v] = (ways[v] + ways[u]) % 1000000007
			}
		}
	}

	return int(ways[n-1])
}

type CountPathsHeap []Item

func (h CountPathsHeap) Len() int           { return len(h) }
func (h CountPathsHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h CountPathsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *CountPathsHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *CountPathsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
