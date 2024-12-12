package solutions

import (
	"container/heap"
	"math"
)

func PickGifts(gifts []int, k int) int64 {
	h := &MaxHeap{}
	heap.Init(h)
	for _, gift := range gifts {
		heap.Push(h, gift)
	}

	for i := k; i > 0; i-- {
		value := heap.Pop(h).(int)
		sq := math.Floor(math.Sqrt(float64(value)))
		heap.Push(h, int(sq))
	}

	ans := 0
	for h.Len() > 0 {
		first := heap.Pop(h).(int)
		ans += first
	}

	return int64(ans)
}
