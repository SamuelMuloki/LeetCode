package solutions

func New21Game(n int, k int, maxPts int) float64 {
	dp := make([]float64, n+1)
	dp[0] = 1.0
	s := 0.0
	if k > 0 {
		s = 1.0
	}

	for i := 1; i <= n; i++ {
		dp[i] = s / float64(maxPts)
		if i < k {
			s += dp[i]
		}
		if i-maxPts >= 0 && i-maxPts < k {
			s -= dp[i-maxPts]
		}
	}

	res := 0.0
	for i := k; i <= n; i++ {
		res += dp[i]
	}

	return res
}
