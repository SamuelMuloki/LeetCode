package solutions

import "sort"

func RangeSum(nums []int, n int, left int, right int) int {
	var arr []int
	arr = append(arr, nums...)
	sum := 0

	for i := 0; i < n; i++ {
		sum = nums[i]
		for j := i + 1; j < n; j++ {
			sum = sum + nums[j]
			arr = append(arr, sum)
		}
	}
	sort.Ints(arr)
	sum = 0
	for i := left - 1; i < right; i++ {
		sum = sum + arr[i]
	}
	return sum % 1000000007
}
