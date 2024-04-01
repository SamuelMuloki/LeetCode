package solutions

func CountSubarrays2(nums []int, minK int, maxK int) int64 {
	var ans int64
	left := 0
	pmin := -1
	pmax := -1

	for right, num := range nums {
		if num < minK || num > maxK {
			left = right + 1
			pmin = -1
			pmax = -1
		} else {
			if num == minK {
				pmin = right
			}
			if num == maxK {
				pmax = right
			}
			ans += int64(max(0, min(pmin, pmax)-left+1))
		}
	}

	return ans
}
