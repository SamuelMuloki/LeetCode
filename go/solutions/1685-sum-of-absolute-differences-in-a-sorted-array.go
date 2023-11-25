package solutions

func GetSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	totalSum := 0
	for i := range nums {
		totalSum += nums[i]
	}

	ans := make([]int, n)
	leftSum := 0
	for i := 0; i < n; i++ {
		rightSum := totalSum - leftSum - nums[i]

		leftCount, rightCount := i, n-1-i

		leftTotal := leftCount*nums[i] - leftSum
		rightTotal := rightSum - rightCount*nums[i]

		ans[i] = leftTotal + rightTotal
		leftSum += nums[i]
	}

	return ans
}
