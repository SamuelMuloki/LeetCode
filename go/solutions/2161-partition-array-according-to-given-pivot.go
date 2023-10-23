package solutions

func PivotArray(nums []int, pivot int) []int {
	var j int
	ans := make([]int, len(nums))
	for i := range nums {
		if nums[i] < pivot {
			ans[j] = nums[i]
			j++
		}
	}

	for i := range nums {
		if nums[i] == pivot {
			ans[j] = nums[i]
			j++
		}
	}

	for i := range nums {
		if nums[i] > pivot {
			ans[j] = nums[i]
			j++
		}
	}

	return ans
}
