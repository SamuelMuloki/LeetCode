package solutions

import "sort"

func MinimumPushes(word string) int {
	count := make([]int, 26)
	for _, letter := range word {
		count[letter-'a']++
	}

	sort.SliceStable(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	ans := 0
	for i, cnt := range count {
		ans += cnt * (i/8 + 1)
	}

	return ans
}
