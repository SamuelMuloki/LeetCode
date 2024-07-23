package solutions

func TimeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	for i := 0; i < len(tickets); i++ {
		if i <= k {
			ans += min(tickets[k], tickets[i])
		} else {
			ans += min(tickets[k]-1, tickets[i])
		}
	}

	return ans
}
