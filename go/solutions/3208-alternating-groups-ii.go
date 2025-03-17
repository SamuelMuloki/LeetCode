package solutions

func NumberOfAlternatingGroupsII(colors []int, k int) int {
	n := len(colors)
	ans, start := 0, 0
	for i := 0; i < n+k; i++ {
		if i > 0 && colors[(i-1)%n] == colors[i%n] {
			start = i % n
		}

		if i-start+1 == k {
			ans++
			start = (start + 1) % n
		}
	}

	return ans
}
