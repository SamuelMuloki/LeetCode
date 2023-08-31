package solutions

func RotatedSearch2(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return true
		}

		// left sorted portion
		if nums[l] < nums[mid] {
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
			// right sorted portion
		} else if nums[l] > nums[mid] {
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			l++
		}
	}

	return false
}
