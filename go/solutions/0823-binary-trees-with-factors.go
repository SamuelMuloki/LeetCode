package solutions

import "sort"

func NumFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	dp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			factor := arr[j]
			quot, rem := arr[i]/factor, arr[i]%factor
			if rem == 0 {
				dp[arr[i]] += dp[factor] * dp[quot]
			}
		}
		dp[arr[i]]++
	}

	ans, mod := 0, 1000000007
	for _, count := range dp {
		ans += count
	}

	return ans % mod
}
