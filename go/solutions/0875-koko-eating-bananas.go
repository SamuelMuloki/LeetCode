package solutions

func MinEatingSpeed(piles []int, h int) int {
	var check = func(mid int) bool {
		cnt := 0
		for i := 0; i < len(piles); i++ {
			cnt += piles[i] / mid
			if piles[i]%mid != 0 {
				cnt++
			}
		}

		return cnt <= h
	}

	max_cnt := 0
	for i := 0; i < len(piles); i++ {
		max_cnt = max(max_cnt, piles[i])
	}

	l, r := 1, max_cnt
	for l < r {
		mid := l + (r-l)/2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
