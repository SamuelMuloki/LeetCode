package solutions

func CountSubstrings(s string) int {
	n, count := len(s), 0
	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		count++
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if string(s[i]) == string(s[i+1]) {
			count++
			dp[i][i+1] = true
		}
	}

	for diff := 2; diff < n; diff++ {
		for i := 0; i < n-diff; i++ {
			j := i + diff
			if string(s[i]) == string(s[j]) && dp[i+1][j-1] {
				count++
				dp[i][j] = true
			}
		}
	}

	return count
}
