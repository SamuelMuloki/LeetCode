package solutions

func ChalkReplacer(chalk []int, k int) int {
	sum := 0
	for i := 0; i < len(chalk); i++ {
		sum += chalk[i]
		if sum > k {
			break
		}
	}

	k = k % sum
	for i := 0; i < len(chalk); i++ {
		if k < chalk[i] {
			return i
		}
		k = k - chalk[i]
	}

	return 0
}
