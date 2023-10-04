package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MaxVowels(s string, k int) int {
	maxV := 0
	dp := make([]int, len(s))

	runes := []rune(s)
	for i := 0; i < k; i++ {
		if !isLowerCaseVowel(runes[i]) {
			continue
		}

		dp[i] = 1
		maxV++
	}

	windowMax := maxV
	for j := k; j < len(runes); j++ {
		if isLowerCaseVowel(runes[j]) {
			dp[j] = 1
		}

		windowMax += dp[j] - dp[j-k]
		maxV = utils.Max(maxV, windowMax)
	}

	return maxV
}

func isLowerCaseVowel(c rune) bool {
	return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u')
}
