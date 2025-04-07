package solutions

func CanPartition(nums []int) bool {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	for i := n - 1; i >= 0; i-- {
		for total := 0; total <= sum; total++ {
			if sum-total == total {
				dp[i][total] = true
				continue
			}

			skip := dp[i+1][total]
			take := false
			if total+nums[i] <= sum {
				take = dp[i+1][total+nums[i]]
			}

			dp[i][total] = skip || take
		}
	}

	return dp[0][0]
}
