// https://leetcode.com/problems/minimum-operations-to-collect-elements/
package solutions

func MinCollectOperations(nums []int, k int) int {
	final, mask, count := (1<<k)-1, 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		count++
		if nums[i] <= k {
			mask |= 1 << (nums[i] - 1)
		}
		if mask == final {
			break
		}
	}

	return count
}
