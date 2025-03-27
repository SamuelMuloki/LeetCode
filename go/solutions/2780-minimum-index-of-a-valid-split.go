package solutions

func MinimumIndex(nums []int) int {
	x, count, n := nums[0], 0, len(nums)
	for _, num := range nums {
		if num == x {
			count++
		} else {
			count--
		}
		if count == 0 {
			x = num
			count = 1
		}
	}

	xCount := 0
	for _, num := range nums {
		if num == x {
			xCount++
		}
	}

	var check = func(i, left, right int) bool {
		return left*2 > i+1 && right*2 > n-i-1
	}

	count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			count++
		}

		remainingCount := xCount - count
		if check(i, count, remainingCount) {
			return i
		}
	}

	return -1
}
