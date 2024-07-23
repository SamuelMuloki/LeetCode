package solutions

func AverageWaitingTime(customers [][]int) float64 {
	nextWaitTime, nextIdleTime := 0, 0
	for i := 0; i < len(customers); i++ {
		nextIdleTime = max(nextIdleTime, customers[i][0]) + customers[i][1]
		nextWaitTime += nextIdleTime - customers[i][0]
	}

	return float64(nextWaitTime) / float64(len(customers))
}
