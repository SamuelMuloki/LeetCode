package solutions

func MaximumCandies(candies []int, k int64) int {
	var check = func(mid int) bool {
		cnt := 0
		for i := 0; i < len(candies); i++ {
			cnt += candies[i] / mid
		}

		return int64(cnt) >= k
	}

	max_cnt := 0
	for i := 0; i < len(candies); i++ {
		max_cnt = max(max_cnt, candies[i])
	}

	l, r := 0, max_cnt
	for l < r {
		mid := (l + r + 1) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}
