package solutions

func PeopleAwareOfSecret(n int, delay int, forget int) int {
	dp := make([]int64, forget)
	dp[0] = 1
	mod := int64(1e9 + 7)
	share := int64(0)
	for i := 1; i < n; i++ {
		idx := i % forget
		addIdx := (i - delay + forget) % forget
		share = (share + dp[addIdx] - dp[idx] + mod) % mod
		dp[idx] = share
	}
	res := int64(0)
	for _, v := range dp {
		res = (res + v) % mod
	}
	return int(res)
}
