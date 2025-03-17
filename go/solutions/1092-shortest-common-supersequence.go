package solutions

func ShortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	i, j := 0, 0
	ans := []byte{}
	for i < m || j < n {
		if i < m && j < n && str1[i] == str2[j] {
			ans = append(ans, str1[i])
			i++
			j++
		} else if i < m && (j == n || dp[i][j] == dp[i+1][j]) {
			ans = append(ans, str1[i])
			i++
		} else {
			ans = append(ans, str2[j])
			j++
		}
	}

	return string(ans)
}
