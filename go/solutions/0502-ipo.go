package solutions

import (
	"container/heap"
	"sort"
)

type Project struct {
	capital int
	profit  int
}

func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]Project, n)

	for i := 0; i < n; i++ {
		projects[i] = Project{capital: capital[i], profit: profits[i]}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	i := 0

	for j := 0; j < k; j++ {
		for i < n && projects[i].capital <= w {
			heap.Push(maxHeap, projects[i].profit)
			i++
		}

		if maxHeap.Len() == 0 {
			break
		}

		w += heap.Pop(maxHeap).(int)
	}

	return w
}
