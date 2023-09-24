package solutions

func LengthOfLIS(nums []int) int {
	dp := make(map[int]int)
	maxLen := 1

	for num := range nums {
		dp[num] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				maxLen = max(maxLen, dp[i])
			}
		}
	}

	return maxLen
}
