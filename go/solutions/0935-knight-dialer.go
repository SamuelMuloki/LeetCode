package solutions

func KnightDialer(n int) int {
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, 10)
	}

	jumps := [][]int{
		{4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0},
		{}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4},
	}

	var dp func(remain, square int) int
	dp = func(remain, square int) int {
		if remain == 0 {
			return 1
		}

		if memo[remain][square] != 0 {
			return memo[remain][square]
		}

		ans := 0
		for _, nextSquare := range jumps[square] {
			ans = (ans + dp(remain-1, nextSquare)) % (1e9 + 7)
		}

		memo[remain][square] = ans
		return ans
	}

	ans := 0
	for square := 0; square < 10; square++ {
		ans = (ans + dp(n-1, square)) % (1e9 + 7)
	}

	return ans
}
