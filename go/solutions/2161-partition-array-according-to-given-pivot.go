package solutions

func PivotArray(nums []int, pivot int) []int {
	l, p, r := 0, 0, len(nums)
	for i := range nums {
		if nums[i] < pivot {
			p++
		} else if nums[i] > pivot {
			r--
		}
	}

	ans := make([]int, len(nums))
	for i := range nums {
		if nums[i] < pivot {
			ans[l] = nums[i]
			l++
		} else if nums[i] == pivot {
			ans[p] = nums[i]
			p++
		} else {
			ans[r] = nums[i]
			r++
		}
	}

	return ans
}
