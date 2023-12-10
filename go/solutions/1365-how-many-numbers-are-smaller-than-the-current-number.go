package solutions

func SmallerNumbersThanCurrent(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				ans[i]++
			}
		}
	}

	return ans
}
