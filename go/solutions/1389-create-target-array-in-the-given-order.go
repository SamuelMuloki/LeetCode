package solutions

func CreateTargetArray(nums []int, index []int) []int {
	ans := make([]int, 0)
	for i := range nums {
		after := ans[index[i]:]
		ans = append([]int{}, ans[:index[i]]...)
		ans = append(ans, nums[i])
		ans = append(ans, after...)
	}

	return ans
}
