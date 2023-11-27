package solutions

func KnightDialer(n int) int {
	prevDp := make([]int, 10)
	for square := 0; square < 10; square++ {
		prevDp[square] = 1
	}

	jumps := [][]int{
		{4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0},
		{}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4},
	}

	for remain := 1; remain < n; remain++ {
		dp := make([]int, 10)
		for square := 0; square < 10; square++ {
			ans := 0
			for _, nextSquare := range jumps[square] {
				ans = (ans + prevDp[nextSquare]) % (1e9 + 7)
			}
			dp[square] = ans
		}
		prevDp = dp
	}

	ans := 0
	for square := 0; square < 10; square++ {
		ans = (ans + prevDp[square]) % (1e9 + 7)
	}

	return ans
}
