package solutions

import (
	"container/heap"
	"math"
)

func ThirdMax(nums []int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	set, maxVal := make(map[int]int), math.MinInt
	for i := range nums {
		if _, ok := set[nums[i]]; ok {
			continue
		}

		set[nums[i]]++
		maxVal = max(maxVal, nums[i])
		heap.Push(minHeap, nums[i])
		if minHeap.Len() > 3 {
			heap.Pop(minHeap)
		}
	}

	if minHeap.Len() == 3 {
		return (*minHeap)[0]
	}

	return maxVal
}
