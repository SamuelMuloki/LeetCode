package solutions

import "container/heap"

func MinimumDifference2(nums []int) int64 {
	n3 := len(nums)
	n := n3 / 3
	part1 := make([]int64, n+1)
	var sum int64 = 0
	ql := &MaxHeap{}
	heap.Init(ql)
	for i := 0; i < n; i++ {
		sum += int64(nums[i])
		heap.Push(ql, nums[i])
	}
	part1[0] = sum
	for i := n; i < n*2; i++ {
		sum += int64(nums[i])
		heap.Push(ql, nums[i])
		sum -= int64(heap.Pop(ql).(int))
		part1[i-(n-1)] = sum
	}

	var part2 int64 = 0
	qr := &MinHeap{}
	heap.Init(qr)
	for i := n*3 - 1; i >= n*2; i-- {
		part2 += int64(nums[i])
		heap.Push(qr, nums[i])
	}
	ans := part1[n] - part2
	for i := n*2 - 1; i >= n; i-- {
		part2 += int64(nums[i])
		heap.Push(qr, nums[i])
		part2 -= int64(heap.Pop(qr).(int))
		if part1[i-n]-part2 < ans {
			ans = part1[i-n] - part2
		}
	}
	return ans
}
