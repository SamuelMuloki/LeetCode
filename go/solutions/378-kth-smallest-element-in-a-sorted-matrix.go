package solutions

import "container/heap"

func KthSmallest2(matrix [][]int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for r := range matrix {
		for c := range matrix[r] {
			heap.Push(h, matrix[r][c])

			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}

	return (*h)[0]
}
