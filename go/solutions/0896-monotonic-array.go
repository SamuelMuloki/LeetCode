package solutions

func IsMonotonic(nums []int) bool {
	st, prev := []int{nums[0]}, 0
	for i := 1; i < len(nums); i++ {
		curr, front := prev, st[len(st)-1]
		if front > nums[i] {
			curr = 1
		} else if front < nums[i] {
			curr = -1
		}

		if prev < 0 && curr > 0 || prev > 0 && curr < 0 {
			return false
		}

		prev = curr
		st = append(st, nums[i])
	}

	return true
}
