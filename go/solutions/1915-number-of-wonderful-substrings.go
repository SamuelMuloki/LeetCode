package solutions

func WonderfulSubstrings(word string) int64 {
	cnt := make([]int64, 1024)
	cnt[0] = 1

	curr := 0
	ans := int64(0)
	for _, c := range word {
		curr ^= 1 << int(c-'a')
		ans += cnt[curr]
		for i := 'a'; i <= 'j'; i++ {
			odd := curr ^ (1 << int(i-'a'))
			ans += cnt[odd]
		}

		cnt[curr]++
	}

	return ans
}
