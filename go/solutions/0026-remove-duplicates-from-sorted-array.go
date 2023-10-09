// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package solutions

func RemoveDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		j := i
		for j < len(nums) && nums[i-1] == nums[j] {
			j++
		}

		if j > i {
			rem := nums[:i]
			if j < len(nums) {
				rem = append(rem, nums[j:]...)
			}
			nums = rem
		}
	}

	return len(nums)
}
