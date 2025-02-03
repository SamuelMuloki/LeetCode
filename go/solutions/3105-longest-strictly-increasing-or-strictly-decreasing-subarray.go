package solutions

func LongestMonotonicSubarray(nums []int) int {
	ans := 1
	inc, dec := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			dec++
			inc = 1
		} else if nums[i-1] < nums[i] {
			inc++
			dec = 1
		} else {
			inc, dec = 1, 1
		}
		ans = max(ans, max(inc, dec))
	}

	return ans
}
