package solutions

import "container/heap"

type SeatManager struct {
	Heap   *MinHeap
	marker int
}

func ReserveConstructor(n int) SeatManager {
	h := &MinHeap{}
	heap.Init(h)
	return SeatManager{h, 0}
}

func (this *SeatManager) Reserve() int {
	if this.Heap.Len() > 0 {
		lowest := (*this.Heap)[0]
		heap.Pop(this.Heap)
		return lowest
	}

	this.marker++
	return this.marker
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.Heap, seatNumber)
}
