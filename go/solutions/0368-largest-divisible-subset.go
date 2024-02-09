package solutions

import "sort"

func LargestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n, maxIndex := len(nums), 0
	gSize, prev := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		gSize[i] = 1
		prev[i] = -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && gSize[j]+1 > gSize[i] {
				gSize[i] = gSize[j] + 1
				prev[i] = j
			}
		}
		if gSize[i] > gSize[maxIndex] {
			maxIndex = i
		}
	}

	var result []int
	for i := maxIndex; i != -1; i = prev[i] {
		result = append([]int{nums[i]}, result...)
	}
	return result
}
