// https://leetcode.com/problems/single-number/
package solutions

func SingleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
