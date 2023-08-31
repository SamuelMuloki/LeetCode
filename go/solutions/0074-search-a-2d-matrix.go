package solutions

func SearchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		nums := matrix[i]
		l, r := 0, len(nums)-1

		for l <= r {
			mid := (l + r) / 2
			if nums[mid] < target {
				l = mid + 1
			} else if nums[mid] > target {
				r = mid - 1
			} else {
				return true
			}
		}
	}

	return false
}
