package solutions

func MaxScoreSightseeingPair(values []int) int {
	n := len(values)
	maxLeftScore := make([]int, n)
	maxLeftScore[0] = values[0]
	maxScore := 0
	for i := 1; i < n; i++ {
		currentRightScore := values[i] - i
		maxScore = max(maxScore, maxLeftScore[i-1]+currentRightScore)
		currentLeftScore := values[i] + i
		maxLeftScore[i] = max(maxLeftScore[i-1], currentLeftScore)
	}

	return maxScore
}
