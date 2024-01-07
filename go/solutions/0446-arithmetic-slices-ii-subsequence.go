package solutions

func NumberOfArithmeticSlices2(nums []int) int {
	ans := 0
	dp := make([]map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			if val, found := dp[j][diff]; found {
				dp[i][diff] += val + 1
				ans += val
			} else {
				dp[i][diff]++
			}
		}
	}

	return ans
}
