package solutions

import "sort"

func LexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	arr := make([]struct {
		Value, Index int
	}, n)

	for i := 0; i < n; i++ {
		arr[i] = struct{ Value, Index int }{Value: nums[i], Index: i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Value < arr[j].Value
	})

	for i := 0; i < n-1; i++ {
		ind := []int{arr[i].Index}
		ptr := i
		for i < n-1 && abs(arr[i].Value-arr[i+1].Value) <= limit {
			i++
			ind = append(ind, arr[i].Index)
		}
		sort.Ints(ind)
		for _, it := range ind {
			nums[it] = arr[ptr].Value
			ptr++
		}
	}
	return nums
}
