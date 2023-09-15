// https://leetcode.com/problems/kth-largest-element-in-a-stream/
package solutions

import (
	"container/heap"
)

type KthLargest struct {
	K    int
	Heap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)

	obj := KthLargest{k, h}
	for _, num := range nums {
		obj.Add(num)
	}

	return obj
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.Heap, val)
	if this.Heap.Len() > this.K {
		heap.Pop(this.Heap)
	}
	return (*this.Heap)[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
