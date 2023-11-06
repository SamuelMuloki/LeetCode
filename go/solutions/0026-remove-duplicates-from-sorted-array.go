// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package solutions

func RemoveDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}

	return len(nums) - count
}
