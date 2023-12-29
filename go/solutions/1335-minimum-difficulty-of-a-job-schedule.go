package solutions

import "math"

func MinDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, d+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(D, L int) int
	dfs = func(D, L int) int {
		if D == 0 && L == n {
			return 0
		}

		if D == 0 || L == n {
			return math.MaxInt
		}

		if memo[L][D] != -1 {
			return memo[L][D]
		}

		currMax := jobDifficulty[L]
		minV := math.MaxInt
		for schedule := L; schedule < n; schedule++ {
			currMax = max(currMax, jobDifficulty[schedule])
			temp := dfs(D-1, schedule+1)
			if temp != math.MaxInt {
				minV = min(minV, temp+currMax)
			}
		}

		memo[L][D] = minV
		return minV
	}

	return dfs(d, 0)
}
