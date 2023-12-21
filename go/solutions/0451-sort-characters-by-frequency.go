package solutions

import "sort"

func FrequencySort(s string) string {
	if len(s) < 2 {
		return s
	}

	freq := [255]int{}
	for i := range s {
		freq[s[i]]++
	}

	ans := []rune(s)
	sort.Slice(ans, func(i, j int) bool {
		if freq[ans[i]] == freq[ans[j]] {
			return ans[i] > ans[j]
		}

		return freq[ans[i]] > freq[ans[j]]
	})

	return string(ans)
}
