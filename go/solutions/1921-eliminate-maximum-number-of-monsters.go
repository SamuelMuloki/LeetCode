package solutions

import "container/heap"

func EliminateMaximum(dist []int, speed []int) int {
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < len(dist); i++ {
		heap.Push(h, (dist[i]-1)/speed[i])
	}

	ans := 0
	for h.Len() > 0 {
		if (*h)[0] < ans {
			break
		}

		ans++
		heap.Pop(h)
	}

	return ans
}
