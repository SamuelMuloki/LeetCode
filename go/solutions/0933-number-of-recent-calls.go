package solutions

import (
	"container/heap"
)

type RecentCounter struct {
	P    int
	Heap *MinHeap
}

func RecentConstructor() RecentCounter {
	h := &MinHeap{}
	heap.Init(h)

	return RecentCounter{0, h}
}

func (this *RecentCounter) Ping(t int) int {
	heap.Push(this.Heap, t)
	for this.Heap.Len() > 0 && ((*this.Heap)[0] < t-3000 || (*this.Heap)[0] > t) {
		heap.Pop(this.Heap)
	}

	return this.Heap.Len()
}
