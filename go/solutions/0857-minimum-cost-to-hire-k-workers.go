package solutions

import (
	"container/heap"
	"math"
	"sort"
)

type Worker struct {
	quality, wage int
}

func MincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([]Worker, len(quality))
	for i := range workers {
		workers[i] = Worker{quality[i], wage[i]}
	}

	sort.Slice(workers, func(i, j int) bool {
		return workers[i].wage*workers[j].quality < workers[j].wage*workers[i].quality
	})

	h, qSum, minCost := make(MaxHeap, 0), 0, math.MaxFloat64
	for _, w := range workers {
		heap.Push(&h, w.quality)
		qSum += w.quality
		if h.Len() > k {
			qSum -= heap.Pop(&h).(int)
		}
		cost := float64(w.wage*qSum) / float64(w.quality)
		if h.Len() == k && cost < minCost {
			minCost = cost
		}
	}

	return minCost
}
