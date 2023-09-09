package solutions

func MinExtraChar(s string, dictionary []string) int {
	set := make(map[string]bool)
	for k := 0; k < len(dictionary); k++ {
		set[dictionary[k]] = true
	}

	n := len(s) + 1
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			sLen := i - j
			if set[s[j:i]] {
				sLen = 0
			}
			dp[i] = min(dp[i], dp[j]+sLen)
		}
	}

	return dp[len(s)]
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
