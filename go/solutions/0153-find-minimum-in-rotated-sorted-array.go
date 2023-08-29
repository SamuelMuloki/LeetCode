// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array
package solutions

func FindMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= nums[0] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
