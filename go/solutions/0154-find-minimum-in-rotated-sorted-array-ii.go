// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii
package solutions

func FindMin2(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}

	return nums[l]
}
