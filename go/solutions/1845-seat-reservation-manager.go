package solutions

import "container/heap"

type SeatManager struct {
	Heap *MinHeap
}

func ReserveConstructor(n int) SeatManager {
	h := &MinHeap{}
	heap.Init(h)

	for i := 1; i <= n; i++ {
		h.Push(i)
	}

	return SeatManager{h}
}

func (this *SeatManager) Reserve() int {
	lowest := (*this.Heap)[0]
	heap.Pop(this.Heap)
	return lowest
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.Heap, seatNumber)
}
