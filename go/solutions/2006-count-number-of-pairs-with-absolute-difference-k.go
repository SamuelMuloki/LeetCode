package solutions

func CountKDifference(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				ans++
			}
		}
	}

	return ans
}
