package solutions

func LargestCombination(candidates []int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		count := 0
		for _, num := range candidates {
			if num&(1<<i) != 0 {
				count++
			}
		}
		ans = max(ans, count)
	}

	return ans
}
