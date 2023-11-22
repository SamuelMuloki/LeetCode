package solutions

import "sort"

func CustomSortString(order string, s string) string {
	letters := [26]int{}
	for i := range order {
		letters[order[i]-'a'] = i + 1
	}

	ans := []rune(s)
	sort.Slice(ans, func(i, j int) bool {
		return letters[ans[i]-'a'] < letters[ans[j]-'a']
	})

	return string(ans)
}
