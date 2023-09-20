package solutions

func MinOperations(nums []int, x int) int {
	target, n := -x, len(nums)
	for _, num := range nums {
		target += num
	}

	if target == 0 {
		return n
	}

	l, maxLength, runningSum := 0, 0, 0

	for r := 0; r < n; r++ {
		runningSum += nums[r]
		for runningSum > target && l <= r {
			runningSum -= nums[l]
			l++
		}

		if runningSum == target {
			currLength := r - l + 1
			if currLength > maxLength {
				maxLength = currLength
			}
		}
	}

	if maxLength != 0 {
		return n - maxLength
	}

	return -1
}
