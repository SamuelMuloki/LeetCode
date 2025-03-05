package solutions

func ColoredCells(n int) int64 {
	ans, cnt := int64(1), int64(0)
	for i := n; i > 0; i-- {
		ans += cnt
		cnt += 4
	}

	return ans
}
