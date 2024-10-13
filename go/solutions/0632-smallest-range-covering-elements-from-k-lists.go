package solutions

import (
	"container/heap"
	"math"
)

type Element struct {
	value, row, col int
}

type MinHeapRange []Element

func (h MinHeapRange) Len() int           { return len(h) }
func (h MinHeapRange) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeapRange) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapRange) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeapRange) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func SmallestRange(nums [][]int) []int {
	MinHeapRange := &MinHeapRange{}
	maxValue := math.MinInt32

	for i := range nums {
		heap.Push(MinHeapRange, Element{nums[i][0], i, 0})
		if nums[i][0] > maxValue {
			maxValue = nums[i][0]
		}
	}

	rangeStart, rangeEnd := 0, math.MaxInt32

	for MinHeapRange.Len() > 0 {
		minElement := heap.Pop(MinHeapRange).(Element)

		if maxValue-minElement.value < rangeEnd-rangeStart {
			rangeStart = minElement.value
			rangeEnd = maxValue
		}

		if minElement.col+1 < len(nums[minElement.row]) {
			nextValue := nums[minElement.row][minElement.col+1]
			heap.Push(MinHeapRange, Element{nextValue, minElement.row, minElement.col + 1})
			if nextValue > maxValue {
				maxValue = nextValue
			}
		} else {
			break // One list is exhausted
		}
	}

	return []int{rangeStart, rangeEnd}
}
