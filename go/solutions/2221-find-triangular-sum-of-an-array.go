package solutions

func TriangularSum(nums []int) int {
	k := len(nums)
	for k > 1 {
		for i := 0; i < k-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		k--
	}

	return nums[0]
}
