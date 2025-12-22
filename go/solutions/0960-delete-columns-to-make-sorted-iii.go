package solutions

func MinDeletionSize3(strs []string) int {
	W := len(strs[0])
	dp := make([]int, W)

	for i := 0; i < W; i++ {
		dp[i] = 1
	}

	for i := W - 2; i >= 0; i-- {
	search:
		for j := i + 1; j < W; j++ {
			for _, row := range strs {
				if row[i] > row[j] {
					continue search
				}
			}
			if dp[i] < 1+dp[j] {
				dp[i] = 1 + dp[j]
			}
		}
	}

	kept := 0
	for _, v := range dp {
		if v > kept {
			kept = v
		}
	}

	return W - kept
}
