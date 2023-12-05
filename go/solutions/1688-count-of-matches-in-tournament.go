package solutions

func NumberOfMatches(n int) int {
	ans := 0
	for teams := n; teams > 1; {
		if teams%2 == 0 {
			ans += (teams / 2)
			teams /= 2
		} else {
			ans += (teams - 1) / 2
			teams = (teams-1)/2 + 1
		}
	}

	return ans
}
