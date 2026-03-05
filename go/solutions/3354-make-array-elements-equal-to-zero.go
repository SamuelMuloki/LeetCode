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
		for j, k := i, i; j >= 0 || k < n; j, k = j-1, k+1 {
			if j >= 0 {
				sum_left += nums[j]
			}
			if k < n {
				sum_right += nums[k]
			}
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
