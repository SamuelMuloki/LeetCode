package solutions

func FindJudge(n int, trust [][]int) int {
	trusted := make([]int, n+1)
	trusts := make([]int, n+1)
	for i := range trust {
		trusted[trust[i][1]]++
		trusts[trust[i][0]]++
	}

	for i := 1; i <= n; i++ {
		if trusted[i] == n-1 && trusts[i] == 0 {
			return i
		}
	}

	return -1
}
