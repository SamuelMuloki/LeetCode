package solutions

func MincostTickets(days []int, costs []int) int {
	isTravelDay := [366]bool{}
	for _, day := range days {
		isTravelDay[day] = true
	}

	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)
	for i := range dp {
		dp[i] = -1
	}

	var solve func(currDay int) int
	solve = func(currDay int) int {
		if currDay > lastDay {
			return 0
		}

		if !isTravelDay[currDay] {
			return solve(currDay + 1)
		}

		if dp[currDay] != -1 {
			return dp[currDay]
		}

		oneDay := costs[0] + solve(currDay+1)
		sevenDay := costs[1] + solve(currDay+7)
		thirtyDay := costs[2] + solve(currDay+30)

		dp[currDay] = min(oneDay, min(sevenDay, thirtyDay))
		return dp[currDay]
	}

	return solve(1)
}
