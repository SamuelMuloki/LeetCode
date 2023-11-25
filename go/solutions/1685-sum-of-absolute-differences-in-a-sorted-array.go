package solutions

func GetSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		leftSum := prefix[i] - nums[i]
		rightSum := prefix[n-1] - prefix[i]

		leftCount, rightCount := i, n-1-i

		leftTotal := leftCount*nums[i] - leftSum
		rightTotal := rightSum - rightCount*nums[i]

		ans[i] = leftTotal + rightTotal
	}

	return ans
}
