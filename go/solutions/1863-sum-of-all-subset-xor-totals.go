package solutions

func SubsetXORSum(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans |= num
	}

	return ans << (len(nums) - 1)
}
