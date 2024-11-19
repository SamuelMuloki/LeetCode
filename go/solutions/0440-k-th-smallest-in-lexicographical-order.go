package solutions

func FindKthNumber(n int, k int) int {
	var countSteps = func(curr, n int) int {
		steps := 0
		first, last := curr, curr
		for first <= n {
			steps += min(n+1, last+1) - first
			first *= 10
			last = last*10 + 9
		}
		return steps
	}

	curr := 1
	k--

	for k > 0 {
		steps := countSteps(curr, n)
		if steps <= k {
			curr++
			k -= steps
		} else {
			curr *= 10
			k--
		}
	}

	return curr
}
