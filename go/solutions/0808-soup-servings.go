package solutions

import "math"

func SoupServings(n int) float64 {
	m := int(math.Ceil(float64(n) / 25.0))
	dp := make(map[[2]int]float64)

	var calculateDP func(i, j int) float64
	calculateDP = func(i, j int) float64 {
		if i <= 0 && j <= 0 {
			return 0.5
		}
		if i <= 0 {
			return 1
		}
		if j <= 0 {
			return 0
		}
		key := [2]int{i, j}
		if val, ok := dp[key]; ok {
			return val
		}
		res := (calculateDP(i-4, j) + calculateDP(i-3, j-1) +
			calculateDP(i-2, j-2) + calculateDP(i-1, j-3)) / 4.0
		dp[key] = res
		return res
	}

	for k := 1; k <= m; k++ {
		if calculateDP(k, k) > 1-1e-5 {
			return 1
		}
	}

	return calculateDP(m, m)
}
