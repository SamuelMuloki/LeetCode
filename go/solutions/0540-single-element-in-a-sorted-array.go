package solutions

func SingleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if (mid&1 == 0 && nums[mid] == nums[mid+1]) || (mid&1 == 1 && nums[mid] == nums[mid-1]) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
