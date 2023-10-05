// https://leetcode.com/problems/majority-element/
package solutions

func MajorityElement(nums []int) int {
	res, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			res = nums[i]
		} else if res == nums[i] {
			count++
		} else {
			count--
		}
	}

	return res
}
