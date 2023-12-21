package solutions

func CreateTargetArray(nums []int, index []int) []int {
	ans := make([]int, len(nums))
	for idx, i := range index {
		for j := len(ans) - 1; j > i; j-- {
			ans[j] = ans[j-1]
		}

		ans[i] = nums[idx]
	}

	return ans
}
