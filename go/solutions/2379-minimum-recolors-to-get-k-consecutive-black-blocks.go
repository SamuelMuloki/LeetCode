package solutions

func MinimumRecolors(blocks string, k int) int {
	start, cnt := 0, 0
	ans := 101
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			cnt++
		}

		if i >= k-1 {
			ans = min(ans, cnt)
			if blocks[start] == 'W' {
				cnt--
			}
			start++
		}
	}

	return ans
}
