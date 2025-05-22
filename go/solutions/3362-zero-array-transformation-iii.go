package solutions

import (
	"container/heap"
	"sort"
)

func MaxRemoval(nums []int, queries [][]int) int {
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})
	pq := &MinHeap{}
	heap.Init(pq)
	diff := make([]int, len(nums)+1)
	operations := 0

	for i, j := 0, 0; i < len(nums); i++ {
		operations += diff[i]
		for j < len(queries) && queries[j][0] == i {
			heap.Push(pq, queries[j][1])
			j++
		}
		for operations < nums[i] && pq.Len() > 0 && (*pq)[0] >= i {
			operations += 1
			diff[heap.Pop(pq).(int)+1] -= 1
		}
		if operations < nums[i] {
			return -1
		}
	}
	return pq.Len()
}
