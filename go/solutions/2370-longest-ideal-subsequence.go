package solutions

func LongestIdealString(s string, k int) int {
	dp := [26]int{}
	ans := 0
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		dp[c] = max(dp[c], 1)

		prevLength := dp[c]
		ans = max(ans, prevLength)
		for j := max(0, c-k); j < min(26, c+k+1); j++ {
			dp[j] = max(dp[j], prevLength+1)
		}
	}
	return ans
}
