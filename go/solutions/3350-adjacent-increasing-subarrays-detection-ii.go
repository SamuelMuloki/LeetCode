package solutions

func MaxIncreasingSubarrays(nums []int) int {
	l, r := 1, len(nums)
	res := 0
	for l <= r {
		mid := l + (r-l)/2
		if HasIncreasingSubarrays(nums, mid) {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
