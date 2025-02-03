package solutions

func CountPartitions(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	ans, currSum := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		currSum += nums[i]
		sum -= nums[i]
		if abs(currSum-sum)%2 == 0 {
			ans++
		}
	}

	return ans
}
