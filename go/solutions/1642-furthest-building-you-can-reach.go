package solutions

import "container/heap"

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	minHeap := &MinHeap{}

	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff <= 0 {
			continue
		}
		heap.Push(minHeap, diff)
		if minHeap.Len() > ladders {
			bricks -= heap.Pop(minHeap).(int)
		}
		if bricks < 0 {
			return i - 1
		}
	}

	return len(heights) - 1
}
