package solutions

func MincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)

	i := 0
	for day := 1; day <= lastDay; day++ {
		if day < days[i] {
			dp[day] = dp[day-1]
		} else {
			i++
			dp[day] = min(
				dp[day-1]+costs[0],
				min(
					dp[max(0, day-7)]+costs[1],
					dp[max(0, day-30)]+costs[2],
				),
			)
		}
	}

	return dp[lastDay]
}
