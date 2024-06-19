package solutions

func MinDays(bloomDay []int, m int, k int) int {
	var getNumOfBouquets = func(mid int) int {
		numOfBouquets, count := 0, 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= mid {
				count++
			} else {
				count = 0
			}

			if count == k {
				numOfBouquets++
				count = 0
			}
		}

		return numOfBouquets
	}

	start, end := 0, 0
	for _, day := range bloomDay {
		end = max(end, day)
	}

	minDays := -1
	for start <= end {
		mid := (start + end) / 2
		if getNumOfBouquets(mid) >= m {
			minDays = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return minDays
}
