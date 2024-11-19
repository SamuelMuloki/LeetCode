package solutions

import (
	"container/heap"
	"sort"
)

func MinGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	pq := &MinHeap{}
	heap.Init(pq)
	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if pq.Len() > 0 && (*pq)[0] < start {
			heap.Pop(pq)
		}
		heap.Push(pq, end)
	}
	return pq.Len()
}
