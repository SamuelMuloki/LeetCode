package solutions

func FindMaxForm(strs []string, m int, n int) int {
	counts := make([][2]int, len(strs))
	for i := 0; i < len(strs); i++ {
		ones, zeroes := 0, 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		counts[i] = [2]int{zeroes, ones}
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, c := range counts {
		zeros, ones := c[0], c[1]
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zeros][j-ones])
			}
		}
	}

	return dp[m][n]
}
