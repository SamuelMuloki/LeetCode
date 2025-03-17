package solutions

import "container/heap"

func MinOperations6(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	ans := 0
	for h.Len() >= 2 {
		top := (*h)[0]
		if top >= k {
			return ans
		}

		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		newNum := min(x, y)*2 + max(x, y)
		heap.Push(h, newNum)
		ans++
	}

	return ans
}
