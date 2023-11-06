package solutions

import "container/heap"

func FindKthLargest(nums []int, k int) int {
	minHeap := MinHeap(nums[:k])
	heap.Init(&minHeap)

	for _, num := range nums[k:] {
		if num > minHeap[0] {
			heap.Pop(&minHeap)
			heap.Push(&minHeap, num)
		}
	}

	return minHeap[0]
}
