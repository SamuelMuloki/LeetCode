package solutions

func MinDeletionSize(strs []string) int {
	n := len(strs[0])
	res := 0
	for j := 0; j < n; j++ {
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] > strs[i][j] {
				res++
				break
			}
		}
	}

	return res
}
