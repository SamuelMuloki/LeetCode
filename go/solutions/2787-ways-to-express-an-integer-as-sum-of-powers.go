package solutions

import "math"

func NumberOfWays3(n int, x int) int {
	mod := 1000000007
	powers := []int{}
	for base := 1; ; base++ {
		pow := int(math.Pow(float64(base), float64(x)))
		if pow > n {
			break
		}
		powers = append(powers, pow)
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for _, p := range powers {
		for t := n; t >= p; t-- {
			dp[t] = (dp[t] + dp[t-p]) % mod
		}
	}

	return dp[n]
}
