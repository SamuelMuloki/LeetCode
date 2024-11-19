package solutions

import (
	"container/heap"
	"math"
)

func MaxKelements(nums []int, k int) int64 {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, num := range nums {
		heap.Push(maxHeap, num)
	}

	ans := 0
	for k > 0 {
		num := heap.Pop(maxHeap).(int)
		ans += num

		ceil := math.Ceil(float64(num) / 3)
		heap.Push(maxHeap, int(ceil))
		k--
	}

	return int64(ans)
}
