package solutions

func MinDeletions(s string) int {
	cnt := make(map[rune]int)
	deletions := 0
	used_frequencies := make(map[int]bool)

	for _, c := range s {
		cnt[c]++
	}

	for _, freq := range cnt {
		for freq > 0 && used_frequencies[freq] {
			freq--
			deletions++
		}
		used_frequencies[freq] = true
	}

	return deletions
}
