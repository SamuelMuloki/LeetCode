package solutions

func MinOperations13(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}

	return res % k
}
