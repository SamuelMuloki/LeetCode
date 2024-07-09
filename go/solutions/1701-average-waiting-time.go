package solutions

func AverageWaitingTime(customers [][]int) float64 {
	n := len(customers)
	lastTime, sum := 0, 0
	for i := 0; i < n; i++ {
		curr := customers[i][0] + customers[i][1]
		if lastTime > customers[i][0] {
			curr = lastTime + customers[i][1]
		}
		sum += curr - customers[i][0]
		lastTime = curr
	}

	return float64(sum) / float64(n)
}
