package solutions

func MaxScoreSightseeingPair(values []int) int {
	n := len(values)
	ans, maxScore := 0, values[0]
	for i := 1; i < n; i++ {
		ans = max(ans, maxScore+values[i]-i)
		maxScore = max(maxScore, values[i]+i)
	}

	return ans
}
