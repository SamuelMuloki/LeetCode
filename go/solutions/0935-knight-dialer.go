package solutions

func KnightDialer(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 10)
	}

	for square := 0; square < 10; square++ {
		dp[0][square] = 1
	}

	jumps := [][]int{
		{4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0},
		{}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4},
	}

	for remain := 1; remain < n; remain++ {
		for square := 0; square < 10; square++ {
			ans := 0
			for _, nextSquare := range jumps[square] {
				ans = (ans + dp[remain-1][nextSquare]) % (1e9 + 7)
			}
			dp[remain][square] = ans
		}
	}

	ans := 0
	for square := 0; square < 10; square++ {
		ans = (ans + dp[n-1][square]) % (1e9 + 7)
	}

	return ans
}
