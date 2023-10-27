package solutions

type Pair struct {
	a, b int
}

func Change(amount int, coins []int) int {
	n := len(coins)
	dp := make(map[Pair]int, n)

	var dfs func(i, val int) int
	dfs = func(i, val int) int {
		pair := Pair{i, val}
		if v, ok := dp[pair]; ok {
			return v
		}

		if val == amount {
			return 1
		}

		if val > amount {
			return 0
		}

		if i == len(coins) {
			return 0
		}

		dp[pair] = dfs(i+1, val) + dfs(i, val+coins[i])
		return dp[pair]
	}

	return dfs(0, 0)
}
