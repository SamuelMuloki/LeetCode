package solutions

func MinimumSteps(s string) int64 {
	totalSwaps := 0
	blackBallCount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			totalSwaps += blackBallCount
		} else {
			blackBallCount++
		}
	}

	return int64(totalSwaps)
}
