package solutions

import "sort"

func NumFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	// map
	// key: root node value
	// value: number of binary tree
	dp := make(map[int]int)
	// scan each possible root node value
	for i := 0; i < len(arr); i++ {
		// Case 1: arr[i] as root with child nodes

		// scan each potential child node value
		for j := 0; j < i; j++ {
			factor := arr[j]
			quot, rem := arr[i]/factor, arr[i]%factor
			// current (factor, quotient) pair are feasible to be child nodes
			if rem == 0 {
				dp[arr[i]] += dp[factor] * dp[quot]
			}
		}
		// Case 2: arr[i] as root without child nodes
		dp[arr[i]]++
	}

	ans, mod := 0, 1000000007
	for _, count := range dp {
		ans += count
	}

	return ans % mod
}
