package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MinDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(word1) && j == len(word2) {
			return 0
		}

		if i == len(word1) || j == len(word2) {
			return utils.Max(len(word1)-i, len(word2)-j)
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		if word1[i] == word2[j] {
			return dfs(i+1, j+1)
		}

		dp[i][j] = 1 + utils.Min(dfs(i+1, j), dfs(i, j+1))
		return dp[i][j]
	}

	return dfs(0, 0)
}
