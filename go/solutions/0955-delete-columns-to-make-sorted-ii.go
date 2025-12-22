package solutions

func MinDeletionSize2(strs []string) int {
	N := len(strs)
	W := len(strs[0])
	ans := 0

	cur := make([]string, N)
	for i := 0; i < N; i++ {
		cur[i] = ""
	}

	for j := 0; j < W; j++ {
		cur2 := make([]string, N)
		copy(cur2, cur)

		for i := 0; i < N; i++ {
			cur2[i] += string(strs[i][j])
		}

		sorted := true
		for i := 0; i < N-1; i++ {
			if cur2[i] > cur2[i+1] {
				sorted = false
				break
			}
		}

		if sorted {
			cur = cur2
		} else {
			ans++
		}
	}

	return ans
}
