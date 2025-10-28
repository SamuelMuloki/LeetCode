package solutions

func CountValidSelections(nums []int) int {
	n := len(nums)
	zeros := make([]int, 0)
	for i, num := range nums {
		if num != 0 {
			continue
		}
		zeros = append(zeros, i)
	}

	res := 0
	for _, i := range zeros {
		sum_left, sum_right := 0, 0
		for j := i; j >= 0; j-- {
			sum_left += nums[j]
		}
		for j := i; j < n; j++ {
			sum_right += nums[j]
		}

		if sum_right == sum_left {
			res += 2
		}

		if sum_right-1 == sum_left ||
			sum_right == sum_left-1 {
			res++
		}

	}
	return res
}
