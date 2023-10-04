// https://leetcode.com/problems/longest-palindromic-substring/
package solutions

func LongestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	output := []int{0, 0}

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if string(s[i]) == string(s[i+1]) {
			dp[i][i+1] = true
			output[0] = i
			output[1] = i + 1
		}
	}

	for diff := 2; diff < n; diff++ {
		for i := 0; i < n-diff; i++ {
			j := i + diff
			if string(s[i]) == string(s[j]) && dp[i+1][j-1] {
				dp[i][j] = true
				output[0] = i
				output[1] = j
			}
		}
	}

	i, j := output[0], output[1]

	return string(s[i : j+1])
}
