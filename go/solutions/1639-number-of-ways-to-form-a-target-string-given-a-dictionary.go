package solutions

func NumWays2(words []string, target string) int {
	n := len(words[0])
	m := len(target)
	freq := make([][26]int, n)

	for _, word := range words {
		for i, c := range word {
			freq[i][c-'a']++
		}
	}

	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[0][j] = freq[j][target[0]-'a']
		if j > 0 {
			dp[0][j] += dp[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = (dp[i][j-1] + dp[i-1][j-1]*freq[j][target[i]-'a']) % 1000000007
		}
	}

	return dp[m-1][n-1]
}
